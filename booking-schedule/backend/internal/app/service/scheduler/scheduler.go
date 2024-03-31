package scheduler

import (
	"log/slog"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/booking"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/rabbit"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	bookingRepository booking.Repository
	log               *slog.Logger
	tracer            trace.Tracer
	rabbitProducer    rabbit.Producer
	checkPeriod       time.Duration
	bookingTTL        time.Duration
}

func NewSchedulerService(bookingRepository booking.Repository, log *slog.Logger, tracer trace.Tracer, rabbitProducer rabbit.Producer, checkPeriod time.Duration, bookingTTL time.Duration) *Service {
	return &Service{
		bookingRepository: bookingRepository,
		log:               log,
		tracer:            tracer,
		rabbitProducer:    rabbitProducer,
		checkPeriod:       checkPeriod,
		bookingTTL:        bookingTTL,
	}
}
