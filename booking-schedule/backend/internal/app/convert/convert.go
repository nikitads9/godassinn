package convert

import (
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/api"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
)

func ToBookingInfo(req *api.Booking) (*model.BookingInfo, error) {
	if req == nil {
		return nil, api.ErrEmptyRequest
	}

	res := &model.BookingInfo{
		ID:        req.BookingID,
		UserID:    req.UserID,
		OfferID:   req.OfferID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if req.NotifyAt.Valid {
		dur, err := time.ParseDuration(req.NotifyAt.String)
		if err != nil {
			return nil, err
		}
		res.NotifyAt = dur
	}

	return res, nil
}

func ToApiBookingInfo(mod *model.BookingInfo) *api.BookingInfo {

	res := &api.BookingInfo{
		ID:        mod.ID,
		OfferID:   mod.OfferID,
		StartDate: mod.StartDate,
		EndDate:   mod.EndDate,
		CreatedAt: mod.CreatedAt,
		UserID:    mod.UserID,
	}

	if mod.NotifyAt != 0 {
		notifyAt := mod.NotifyAt.String()
		res.NotifyAt = &notifyAt
	}

	if mod.UpdatedAt.Valid {
		res.UpdatedAt = &mod.UpdatedAt.Time
	}

	return res
}

func ToApiBookingsInfo(bookings []*model.BookingInfo) []*api.BookingInfo {
	if bookings == nil {
		return nil
	}

	res := make([]*api.BookingInfo, 0, len(bookings))
	for _, elem := range bookings {
		res = append(res, ToApiBookingInfo(elem))
	}

	return res
}

func ToApiOffers(mod []*model.Offer) []*api.Offer {
	var res []*api.Offer
	for _, elem := range mod {
		res = append(res, &api.Offer{
			OfferID:          elem.OfferID,
			Name:             elem.Name,
			Cost:             elem.Cost,
			City:             elem.City,
			Street:           elem.Street,
			House:            elem.House,
			Rating:           elem.Rating,
			TypeOfHousing:    elem.TypeOfHousing,
			BedsCount:        elem.BedsCount,
			ShortDescription: elem.ShortDescription,
			LandlordID:       elem.LandlordID,
		})
	}

	return res
}

// Эта функция преобразует массив занятых интервалов к виду свободных
func ToVacantDates(mod []*model.Interval) []*api.Interval {
	now := time.Now()
	month := now.Add(720 * time.Hour)
	var res []*api.Interval

	if mod == nil {
		res = append(res, &api.Interval{
			StartDate: now,
			EndDate:   month,
		})
		return res
	}

	if now.Before(mod[0].StartDate) {
		res = append(res, &api.Interval{
			StartDate: now,
			EndDate:   mod[0].StartDate,
		})
	}

	if len(mod) == 1 && mod[0].EndDate.After(month) {
		return res
	}

	if len(mod) == 1 {
		res = append(res, &api.Interval{
			StartDate: mod[0].EndDate,
			EndDate:   month,
		})
		return res
	}

	for i := 1; i < len(mod); i++ {
		if mod[i].EndDate.Before(month) {
			res = append(res, &api.Interval{
				StartDate: mod[i-1].EndDate,
				EndDate:   mod[i].StartDate,
			})
		} else {
			res = append(res, &api.Interval{
				StartDate: mod[i-1].EndDate,
				EndDate:   mod[i].StartDate,
			})
			return res
		}

	}

	if mod[len(mod)-1].EndDate.Before(month) {
		res = append(res, &api.Interval{
			StartDate: mod[len(mod)-1].EndDate,
			EndDate:   month,
		})
	}

	return res
}

func ToUserInfo(user *api.SignUpRequest) (*model.User, error) {
	if user == nil {
		return nil, api.ErrEmptyRequest
	}

	mod := &model.User{
		Login:       user.Login,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		CreatedAt:   time.Now(),
	}
	return mod, nil
}

func ToApiUserInfo(user *model.User) *api.UserInfo {
	res := &api.UserInfo{
		ID:          user.ID,
		Login:       user.Login,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   time.Now(),
	}

	if user.UpdatedAt != nil {
		res.UpdatedAt = user.UpdatedAt
	}

	return res
}

func ToUpdateUserInfo(user *api.EditMyProfileRequest, userID int64) *model.UpdateUserInfo {
	mod := &model.UpdateUserInfo{
		ID:          userID,
		Login:       user.Login,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}

	return mod
}
