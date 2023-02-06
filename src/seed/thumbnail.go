package main

import (
	"app/src/db"
	"app/src/model"

	"fmt"
)

var thumbCountByHotel = 9

func generateSeedThumb(hotelId string) *model.Thumbnail {
	return &model.Thumbnail{
		ID:      UuidV4(),
		HotelID: hotelId,
	}
}

func seedThumbs(db *db.DB, hotels []*model.Hotel) []*model.Thumbnail {
	index := 0
	thumbs := []*model.Thumbnail{}
	// total: 45, 各ホテルに9枚ずつ
	for hotelIndex := 0; hotelIndex < hotelTotalCount; hotelIndex++ {
		for thumbIndex := 0; thumbIndex < thumbCountByHotel; thumbIndex++ {
			hotelID := hotels[hotelIndex].ID
			t := generateSeedThumb(hotelID)
			if err := db.Conn.Create(&t).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			if err := db.Conn.Last(&t).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			thumbs = append(thumbs, t)
			index++
		}
	}
	return thumbs
}
