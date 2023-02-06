package model

import (
	"time"
)

type Vacancy struct {
	ID        string `gorm:"type:uuid;primary_key;not_null;"`
	Empty     bool
	HotelID   string
	CreatedAt time.Time `gorm:"type:time;"`
	UpdatedAt time.Time `gorm:"type:time;"`
}
