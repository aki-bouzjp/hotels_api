package model

import (
	"time"
)

type Room struct {
	ID        string `gorm:"type:uuid;primary_key;not_null;"`
	Name      string
	GroupName string
	HotelID   string
	CreatedAt time.Time `gorm:"type:time;"`
	UpdatedAt time.Time `gorm:"type:time;"`

	Images []Image `gorm:"many2many:room_images";`
}
