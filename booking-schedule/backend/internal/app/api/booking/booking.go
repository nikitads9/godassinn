package booking

import (
	"errors"
	"net/http"

	bookingRepo "github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/booking"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/service/booking"

	"go.opentelemetry.io/otel/trace"
)

type Implementation struct {
	booking *booking.Service
	tracer  trace.Tracer
}

var (
	errNoBookingID = errors.New("received no booking id")
	errNoInterval  = errors.New("received no time period")
	errNoSuiteID   = errors.New("received no suite id")
	//ErrBookingNotFound = errors.New("no booking with this id")
)

func NewImplementation(booking *booking.Service, tracer trace.Tracer) *Implementation {
	return &Implementation{
		booking: booking,
		tracer:  tracer,
	}
}

func GetErrorCode(err error) int {
	switch err {
	case bookingRepo.ErrNotFound:
		return http.StatusNotFound
	case bookingRepo.ErrNoRowsAffected:
		return http.StatusNotFound
	case bookingRepo.ErrUnauthorized:
		return http.StatusUnauthorized
	case booking.ErrNotAvailible:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
