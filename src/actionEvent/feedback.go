package actionEvent

import "time"

type Feedback struct {
	Uuid      string    `csv:"uuid"`
	HotelID   string    `csv:"hotel_id"`
	Action    string    `csv:"action"`
	ZoomLevel int       `csv:"zoom_level"`
	CreatedAt time.Time `csv:"created_at"`
}
