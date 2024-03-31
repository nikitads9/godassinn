package user

import (
	"context"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func (s *Service) EditUser(ctx context.Context, user *model.UpdateUserInfo) error {
	return s.userRepository.EditUser(ctx, user)
}
