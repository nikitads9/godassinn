package booking

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/db"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"go.opentelemetry.io/otel/trace"
)

type Repository interface {
	AddBooking(ctx context.Context, mod *model.BookingInfo) (uuid.UUID, error)
	GetBooking(ctx context.Context, bookingID uuid.UUID, userID int64) (*model.BookingInfo, error)
	GetBookings(ctx context.Context, startDate time.Time, endDate time.Time, userID int64) ([]*model.BookingInfo, error)
	UpdateBooking(ctx context.Context, mod *model.BookingInfo) error
	DeleteBooking(ctx context.Context, bookingID uuid.UUID, userID int64) error
	GetVacantOffers(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Offer, error)
	GetBusyDates(ctx context.Context, offerID int64) ([]*model.Interval, error)
	GetBookingListByDate(ctx context.Context, start time.Time, end time.Time) ([]*model.BookingInfo, error)
	DeleteBookingsBeforeDate(ctx context.Context, end time.Time) error
	CheckAvailibility(ctx context.Context, mod *model.BookingInfo) (*model.Availibility, error)
}

var (
	ErrNotFound       = errors.New("no booking with this id")
	ErrNoRowsAffected = errors.New("no database entries affected by this operation")
	ErrUnauthorized   = errors.New("no user associated with this token")

	ErrQuery        = errors.New("failed to execute query")
	ErrQueryBuild   = errors.New("failed to build query")
	ErrPgxScan      = errors.New("failed to read database response")
	ErrNoConnection = errors.New("could not connect to database")
	ErrUuid         = errors.New("failed to generate uuid")

	pgNoConnection = new(*pgconn.ConnectError)
	ErrNoSuchUser  = "ERROR: violates foreign key constraint \"fk_users\" (SQLSTATE 23503)"
)

type repository struct {
	client db.Client
	log    *slog.Logger
	tracer trace.Tracer
}

func NewBookingRepository(client db.Client, log *slog.Logger, tracer trace.Tracer) Repository {
	return &repository{
		client: client,
		log:    log,
		tracer: tracer,
	}
}
