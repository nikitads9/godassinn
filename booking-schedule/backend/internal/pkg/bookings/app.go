package bookings

import (
	"context"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/api"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/middleware/auth"
	mwLogger "github.com/nikitads9/godassinn/booking-schedule/backend/internal/middleware/logger"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/middleware/metrics"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/certificates"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/observability"

	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/metric"

	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/riandyrn/otelchi"
)

type App struct {
	configType string
	pathConfig string

	pathCert string
	pathKey  string

	serviceProvider *serviceProvider
	router          *chi.Mux
	meter           metric.Meter
}

// NewApp ...
func NewApp(ctx context.Context, configType string, pathConfig string, pathCert string, pathKey string) (*App, error) {
	a := &App{
		configType: configType,
		pathConfig: pathConfig,
		pathCert:   pathCert,
		pathKey:    pathKey,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {

	inits := []func(context.Context) error{
		a.initMeter,
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.configType, a.pathConfig, a.meter)

	return nil
}

// Run ...
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close() //nolint:errcheck
	}()

	err := a.startServer()
	if err != nil {
		a.serviceProvider.GetLogger().Error("failed to start server: %s", err)
		return err
	}

	return nil
}

func (a *App) initMeter(ctx context.Context) error {
	meter, err := observability.NewMeter(ctx, "booking")
	if err != nil {
		return err
	}

	a.meter = meter

	return nil
}

func (a *App) startServer() error {
	if a.meter != nil {
		go observability.CollectMachineResourceMetrics(a.meter, a.serviceProvider.GetLogger())
	}

	srv := a.serviceProvider.getServer(a.router)
	if srv == nil {
		a.serviceProvider.GetLogger().Error("server was not initialized")
		return errors.New("server was not initialized")
	}
	a.serviceProvider.GetLogger().Info("starting server", slog.String("address", srv.Addr))

	done := make(chan os.Signal, 1)
	errChan := make(chan error)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(errChan)
		switch a.serviceProvider.GetConfig().GetEnv() {
		case envProd:
			if _, err := os.Stat(a.pathKey); err != nil {
				err := certificates.InitCertificates(a.pathCert, a.pathKey)
				if err != nil {
					a.serviceProvider.GetLogger().Error("failed to initialize certificates", sl.Err(err))
					errChan <- err
				}
			}

			if err := srv.ListenAndServeTLS(a.pathCert, a.pathKey); err != nil {
				a.serviceProvider.GetLogger().Error("", sl.Err(err))
				errChan <- err
			}
		default:
			if err := srv.ListenAndServe(); err != nil {
				a.serviceProvider.GetLogger().Error("", sl.Err(err))
				errChan <- err
			}
		}
	}()

	a.serviceProvider.GetLogger().Info("server started")

	select {
	case err := <-errChan:
		return err
	case <-done:
		a.serviceProvider.GetLogger().Info("stopping server")
		// TODO: move timeout to config
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			a.serviceProvider.GetLogger().Error("failed to stop server", sl.Err(err))
			return err
		}

		a.serviceProvider.GetLogger().Info("server stopped")
	}

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	bookingImpl := a.serviceProvider.GetBookingImpl(ctx)
	userImpl := a.serviceProvider.GetUserImpl(ctx)

	address, err := a.serviceProvider.GetConfig().GetAddress()
	if err != nil {
		return err
	}
	a.serviceProvider.GetLogger().Info("initializing server", slog.String("address", address))
	a.serviceProvider.GetLogger().Debug("logger debug mode enabled")

	a.router = chi.NewRouter()

	a.router.Group(func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{a.serviceProvider.GetConfig().GetTracerConfig().EndpointURL},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}, //TODO allow only real otlp methods
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
		}))
		r.Handle("/metrics", promhttp.Handler())
	})

	a.router.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(otelchi.Middleware("bookings-api", otelchi.WithChiRoutes(a.router)))
		r.Use(metrics.NewMetricMiddleware(a.serviceProvider.GetMeter(ctx)))
		r.Use(mwLogger.New(a.serviceProvider.GetLogger()))
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"}, //"X-CSRF-Token" for tokens stored in cookies
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
		r.Use(middleware.Recoverer)
		r.Route("/bookings", func(r chi.Router) {
			r.Get("/ping", api.HandlePingCheck())
			r.Route("/user", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(auth.Auth(a.serviceProvider.GetLogger(), a.serviceProvider.GetJWTService(ctx)))
					r.Get("/me", userImpl.GetMyProfile(a.serviceProvider.GetLogger()))
					r.Delete("/delete", userImpl.DeleteMyProfile(a.serviceProvider.GetLogger()))
					r.Patch("/edit", userImpl.EditMyProfile(a.serviceProvider.GetLogger()))
				})

			})
			r.Get("/get-vacant-rooms", bookingImpl.GetVacantRooms(a.serviceProvider.GetLogger()))
			r.Get("/{suite_id}/get-vacant-dates", bookingImpl.GetVacantDates(a.serviceProvider.GetLogger()))
			r.Group(func(r chi.Router) {
				r.Use(auth.Auth(a.serviceProvider.GetLogger(), a.serviceProvider.GetJWTService(ctx)))
				r.Post("/add", bookingImpl.AddBooking(a.serviceProvider.GetLogger()))
				r.Get("/get-bookings", bookingImpl.GetBookings(a.serviceProvider.GetLogger()))
				r.Route("/{booking_id}", func(r chi.Router) {
					r.Get("/get", bookingImpl.GetBooking(a.serviceProvider.GetLogger()))
					r.Patch("/update", bookingImpl.UpdateBooking(a.serviceProvider.GetLogger()))
					r.Delete("/delete", bookingImpl.DeleteBooking(a.serviceProvider.GetLogger()))
				})
			})

		})
	})

	return nil
}
