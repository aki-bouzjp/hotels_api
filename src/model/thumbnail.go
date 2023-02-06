package model

import (
	"time"
)

type Thumbnail struct {
	ID        string `gorm:"type:uuid;primary_key;not_null;"`
	HotelID   string
	CreatedAt time.Time `gorm:"type:time;"`
	UpdatedAt time.Time `gorm:"type:time;"`

	Images []Image `gorm:"many2many:thumbnail_images";`
}
