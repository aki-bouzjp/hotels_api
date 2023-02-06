package model

import (
	"time"
)

type RoomImage struct {
	ID        string `gorm:"type:uuid;primary_key;not_null;"`
	Size      string
	RoomID    string
	ImageID   string
	CreatedAt time.Time `gorm:"type:time;"`
	UpdatedAt time.Time `gorm:"type:time;"`
}
