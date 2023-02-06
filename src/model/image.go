package model

import (
	"time"
)

type Image struct {
	ID        string `gorm:"type:uuid;primary_key;not_null;"`
	MimeType  string
	Url       string
	Width     int
	Height    int
	CreatedAt time.Time `gorm:"type:time;"`
	UpdatedAt time.Time `gorm:"type:time;"`
}
