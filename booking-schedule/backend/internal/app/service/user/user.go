package user

import (
	"errors"
	"log/slog"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/user"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/service/jwt"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	userRepository user.Repository
	jwtService     jwt.Service
	log            *slog.Logger
	tracer         trace.Tracer
}

var (
	ErrBadLogin  = errors.New("incorrect nickname")
	ErrBadPasswd = errors.New("incorrect password")

	ErrHashFailed = errors.New("failed to hash password")
)

func NewUserService(userRepository user.Repository, jwtService jwt.Service, log *slog.Logger, tracer trace.Tracer) *Service {
	return &Service{
		userRepository: userRepository,
		jwtService:     jwtService,
		log:            log,
		tracer:         tracer,
	}
}
