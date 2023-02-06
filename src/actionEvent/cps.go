package actionEvent

type Cps struct {
	Uuid      string  `json:"uuid"`
	HotelId   string  `json:"hotel_id"`
	Lng       float64 `json:"lng"`
	Lat       float64 `json:"lat"`
	ZoomLevel int     `json:"zoom_level"`
}
