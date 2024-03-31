package booking

import (
	"context"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func (s *Service) GetVacantRooms(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Suite, error) {
	return s.bookingRepository.GetVacantRooms(ctx, startDate, endDate)
}
