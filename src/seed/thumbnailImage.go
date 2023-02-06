package main

import (
	"app/src/db"
	"app/src/model"

	"fmt"
)

func generateSeedThumbImage(size string, thumbId string, imageId string) *model.ThumbnailImage {
	return &model.ThumbnailImage{
		ID:          UuidV4(),
		Size:        size,
		ThumbnailID: thumbId,
		ImageID:     imageId,
	}
}

func seedThumbImages(db *db.DB, thumbs []*model.Thumbnail, images []*model.Image) []*model.ThumbnailImage {
	index := 0
	imageIndex := 0
	thumbImages := []*model.ThumbnailImage{}
	sizes := []string{"original", "small", "medium", "large"}
	// total: 135 = 45 x 3, 135枚（5ホテル x 9枚サムネ）に3枚ずつ各サイズの画像
	for hotelIndex := 0; hotelIndex < hotelTotalCount; hotelIndex++ {
		for thumbIndex := 0; thumbIndex < thumbCountByHotel; thumbIndex++ {
			for _, s := range sizes {
				thumbID := thumbs[index].ID
				imageID := images[imageIndex].ID
				t := generateSeedThumbImage(s, thumbID, imageID)
				if err := db.Conn.Create(&t).Error; err != nil {
					fmt.Printf("%+v", err)
					continue
				}
				if err := db.Conn.Last(&t).Error; err != nil {
					fmt.Printf("%+v", err)
					continue
				}
				thumbImages = append(thumbImages, t)
				imageIndex++
			}
			index++
		}
	}
	return thumbImages
}
