package booking

import (
	"context"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"

	"github.com/gofrs/uuid"
)

func (s *Service) GetBooking(ctx context.Context, bookingID uuid.UUID, userID int64) (*model.BookingInfo, error) {
	return s.bookingRepository.GetBooking(ctx, bookingID, userID)
}
