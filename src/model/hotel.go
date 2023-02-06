package model

import (
	"time"
)

type Hotel struct {
	ID                     string `gorm:"type:uuid;primary_key;not_null;"`
	HappyHotelID           string
	HotelType              string
	Lng                    float64
	Lat                    float64
	Name                   string
	Address                string
	Url                    string
	Homepage               string
	Tel                    string
	Pr                     string
	Score                  float32
	SinglePrice            float32
	RestMinPrice           float32
	LodgingMinPrice        float32
	TransferTimeByWalk     string
	TransferTimeByDrive    string
	RoomCount              int
	ParkingCount           int
	ParkingCountByHighroof int
	ParkingDescription     string
	Mapcode                string
	PricingPlans           string
	OnlineBooking          bool
	OnlineChecking         bool
	CreatedAt              time.Time `gorm:"type:time;"`
	UpdatedAt              time.Time `gorm:"type:time;"`

	Thumbnails []Thumbnail `gorm:"foreginkey:HotelID"`
	Vacancy    Vacancy     `gorm:"foreignkey:HotelID"`
	Rooms      []Room      `gorm:"foreginkey:HotelID"`
}
