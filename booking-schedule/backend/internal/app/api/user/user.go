package user

import (
	"net/http"

	userRepo "github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/user"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/service/user"

	"go.opentelemetry.io/otel/trace"
)

type Implementation struct {
	user   *user.Service
	tracer trace.Tracer
}

func NewImplementation(user *user.Service, tracer trace.Tracer) *Implementation {
	return &Implementation{
		user:   user,
		tracer: tracer,
	}
}

func GetErrorCode(err error) int {
	switch err {
	case userRepo.ErrNoRowsAffected:
		return http.StatusNotFound
	case userRepo.ErrNotFound:
		return http.StatusNotFound
	case userRepo.ErrAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
