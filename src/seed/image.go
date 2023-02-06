package main

import (
	"app/src/db"
	"app/src/model"

	"fmt"
)

func generateSeedImage() *model.Image {
	return &model.Image{
		ID:       UuidV4(),
		MimeType: "image/png",
		Url:      "http://test.co.jp/image.png",
		Width:    100,
		Height:   100,
	}
}

func seedImages(db *db.DB) []*model.Image {
	images := []*model.Image{}
	// total 135*2
	for imageIndex := 1; imageIndex <= 200; imageIndex++ {
		i := generateSeedImage()
		if err := db.Conn.Create(&i).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
		if err := db.Conn.Last(&i).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
		images = append(images, i)
	}
	return images
}
