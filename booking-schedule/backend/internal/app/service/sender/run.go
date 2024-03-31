package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"

	"github.com/streadway/amqp"
)

func (s *Service) Run(ctx context.Context) {
	const op = "service.sender.Run"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("sender service initiated")

	msgChan, err := s.rabbitConsumer.Consume()
	if err != nil {
		log.Error("could not get channel to receive messages: ", sl.Err(err))
		os.Exit(1)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgChan:
			err = s.receiveBookings(msg)
			if err != nil {
				log.Error("could not receive message: ", sl.Err(err))
			}
			err = msg.Ack(false)
			if err != nil {
				log.Error("could not acknowledge message acquiring: ", sl.Err(err))
			}
		}

	}

}

func (s *Service) receiveBookings(msg amqp.Delivery) error {
	const op = "service.sender.receiveBookings"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info(fmt.Sprintf("Received a message: %s", msg.Body))

	var booking *model.BookingInfo
	err := json.Unmarshal(msg.Body, &booking)
	if err != nil {
		log.Error("failed to unmarshal message", sl.Err(err))
		return err
	}

	log.Info(fmt.Sprintf(
		"Booking:  %d \n "+
			"SuiteID: %d \n "+
			"StartDate: %v \n "+
			"EndDate: :%v \n "+
			"NotifyAt: %v \n "+
			"OwnerID: %d \n "+
			"CreatedAt: %v \n "+
			"UpdatedAt: %v \n\n ",
		booking.ID,
		booking.SuiteID,
		booking.StartDate,
		booking.EndDate,
		booking.NotifyAt,
		booking.UserID,
		booking.CreatedAt,
		booking.UpdatedAt,
	))

	return nil
}
