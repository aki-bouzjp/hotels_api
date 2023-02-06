package actionEvent

import "time"

type Session struct {
	Uuid      string    `csv:"uuid"`
	Lng       float64   `csv:"lng"`
	Lat       float64   `csv:"lat"`
	ZoomLevel int       `csv:"zoom_level"`
	CreatedAt time.Time `csv:"created_at"`
}
