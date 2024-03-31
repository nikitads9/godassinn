package booking

import (
	"context"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func (s *Service) GetBookings(ctx context.Context, startDate time.Time, endDate time.Time, id int64) ([]*model.BookingInfo, error) {
	return s.bookingRepository.GetBookings(ctx, startDate, endDate, id)
}
