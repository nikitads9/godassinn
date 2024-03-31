package sender

import (
	"log/slog"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/rabbit"
)

type Service struct {
	log            *slog.Logger
	rabbitConsumer rabbit.Consumer
}

func NewSenderService(log *slog.Logger, rabbitConsumer rabbit.Consumer) *Service {
	return &Service{
		log:            log,
		rabbitConsumer: rabbitConsumer,
	}
}
