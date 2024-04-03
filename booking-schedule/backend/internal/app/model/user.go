package model

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID          int64      `db:"id"`
	Login       string     `db:"login"`
	Name        string     `db:"name"`
	PhoneNumber string     `db:"phone_number"`
	Password    string     `db:"password"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type UpdateUserInfo struct {
	ID          int64       `db:"id"`
	Login       null.String `db:"login"`
	Name        null.String `db:"name"`
	PhoneNumber null.String `db:"phone_number"`
	Password    null.String `db:"password"`
}
