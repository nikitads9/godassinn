package booking

import (
	"context"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func (s *Service) GetVacantOffers(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Offer, error) {
	return s.bookingRepository.GetVacantOffers(ctx, startDate, endDate)
}
