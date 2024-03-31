package auth

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
	case user.ErrBadLogin:
		return http.StatusUnauthorized
	case user.ErrBadPasswd:
		return http.StatusUnauthorized
	case userRepo.ErrNotFound:
		return http.StatusNotFound
	case userRepo.ErrAlreadyExists:
		return http.StatusBadRequest
	case userRepo.ErrDuplicate:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
