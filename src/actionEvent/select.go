package actionEvent

import "time"

type Select struct {
	Uuid      string    `csv:"uuid"`
	HotelID   string    `csv:"hotel_id"`
	Lng       float64   `csv:"lng"`
	Lat       float64   `csv:"lat"`
	ZoomLevel int       `csv:"zoom_level"`
	CreatedAt time.Time `csv:"created_at"`
}
