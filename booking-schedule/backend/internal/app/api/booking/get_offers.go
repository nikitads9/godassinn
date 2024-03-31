package booking

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/api"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/convert"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// GetVacantOffers godoc
//
//	@Summary		Get list of vacant offers
//	@Description	Receives two dates as query parameters. start is to be before end and both should not be expired. Responds with list of vacant offers and their parameters for given interval.
//	@ID				getOffersByDates
//	@Tags			bookings
//	@Produce		json
//	@Param			start	query	string	true	"start"	Format(time.Time) default(2024-03-28T17:43:00)
//	@Param			end	query	string	true	"end"	Format(time.Time) default(2024-03-29T17:43:00)
//	@Success		200	{object}	api.GetVacantOffersResponse
//	@Failure		400	{object}	api.errResponse
//	@Failure		404	{object}	api.errResponse
//	@Failure		503	{object}	api.errResponse
//	@Router			/get-vacant-offers [get]
func (i *Implementation) GetVacantOffers(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "api.booking.GetVacantOffers"

		ctx := r.Context()
		requestID := middleware.GetReqID(ctx)

		log := logger.With(
			slog.String("op", op),
			slog.String("request_id", requestID),
		)
		ctx, span := i.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
		defer span.End()

		start := r.URL.Query().Get("start")
		if start == "" {
			span.RecordError(errNoInterval)
			span.SetStatus(codes.Error, errNoInterval.Error())
			log.Error("invalid request", sl.Err(errNoInterval))
			api.WriteWithError(w, http.StatusBadRequest, errNoInterval.Error())
			return
		}

		span.AddEvent("startDate extracted from query", trace.WithAttributes(attribute.String("start", start)))

		end := r.URL.Query().Get("end")
		if end == "" {
			span.RecordError(errNoInterval)
			span.SetStatus(codes.Error, errNoInterval.Error())
			log.Error("invalid request", sl.Err(errNoInterval))
			api.WriteWithError(w, http.StatusBadRequest, errNoInterval.Error())
			return
		}

		span.AddEvent("endDate extracted from query", trace.WithAttributes(attribute.String("end", end)))

		startDate, err := time.Parse("2006-01-02T15:04:05", start)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("invalid request", sl.Err(err))
			api.WriteWithError(w, http.StatusBadRequest, api.ErrParse.Error())
			return
		}
		endDate, err := time.Parse("2006-01-02T15:04:05", end)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("invalid request", sl.Err(err))
			api.WriteWithError(w, http.StatusBadRequest, api.ErrParse.Error())
			return
		}

		span.AddEvent("start and end dates parsed")

		err = api.CheckDates(startDate, endDate)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("invalid request", sl.Err(err))
			api.WriteWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		span.AddEvent("dates verified")

		offers, err := i.booking.GetVacantOffers(ctx, startDate, endDate)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("internal error", sl.Err(err))
			api.WriteWithError(w, GetErrorCode(err), err.Error())
			return
		}

		span.AddEvent("vacant offers acquired", trace.WithAttributes(attribute.Int("quantity", len(offers))))
		log.Info("vacant offers acquired", slog.Int("quantity: ", len(offers)))

		render.Status(r, http.StatusCreated)
		api.WriteWithStatus(w, http.StatusOK, api.GetVacantOffersResponse{
			Offers: convert.ToApiOffers(offers),
		})
	}
}
