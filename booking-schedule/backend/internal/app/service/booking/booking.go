package booking

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/booking"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/service/jwt"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/db"

	"github.com/jackc/pgx/v5/pgconn"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	bookingRepository booking.Repository
	jwtService        jwt.Service
	log               *slog.Logger
	tracer            trace.Tracer
	txManager         db.TxManager
}

var (
	ErrNotAvailible = errors.New("this period is not availible for booking")

	ErrNoConnection = errors.New("can't begin transaction, no connection to database")
	pgNoConnection  = new(*pgconn.ConnectError)
)

func NewBookingService(bookingRepository booking.Repository, jwtService jwt.Service, log *slog.Logger, txManager db.TxManager, tracer trace.Tracer) *Service {
	return &Service{
		bookingRepository: bookingRepository,
		jwtService:        jwtService,
		log:               log,
		tracer:            tracer,
		txManager:         txManager,
	}
}
