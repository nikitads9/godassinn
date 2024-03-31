package user

import (
	"context"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func (s *Service) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	return s.userRepository.GetUser(ctx, userID)
}
