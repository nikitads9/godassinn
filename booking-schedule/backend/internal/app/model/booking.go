package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v3"
)

type BookingInfo struct {
	ID        uuid.UUID     `db:"id"`
	OfferID   int64         `db:"suite_id"`
	StartDate time.Time     `db:"start_date"`
	EndDate   time.Time     `db:"end_date"`
	NotifyAt  time.Duration `db:"notify_at"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt null.Time     `db:"updated_at"`
	UserID    int64         `db:"user_id"`
}

type Interval struct {
	StartDate time.Time `db:"start"`
	EndDate   time.Time `db:"end"`
}

type Offer struct {
	OfferID          int64  `db:"offer_id"`
	Name             string `db:"name"`
	Cost             int64  `db:"cost"`
	City             string `db:"city"`
	Street           string `db:"street"`
	House            int64  `db:"house"`
	Rating           int64  `db:"rating"`
	Type             string `db:"type"`
	BedsCount        uint8  `db:"beds_count"`
	ShortDescription string `db:"short_description"`
}

type Availibility struct {
	Availible        bool `db:"availible"`
	OccupiedByClient bool `db:"occupied_by_client"`
}
