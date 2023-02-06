package main

import (
	"app/src/db"
	"app/src/model"

	"fmt"
)

func generateSeedRoomImage(size string, roomId string, imageId string) *model.RoomImage {
	return &model.RoomImage{
		ID:      UuidV4(),
		Size:    size,
		RoomID:  roomId,
		ImageID: imageId,
	}
}

func seedRoomImages(db *db.DB, rooms []*model.Room, images []*model.Image) []*model.RoomImage {
	index := 0
	roomImages := []*model.RoomImage{}
	sizes := []string{"original", "small", "medium", "large"}
	// total: 135 = 45 x 3, 45ルームに3枚ずつ
	for roomIndex := 0; roomIndex < hotelTotalCount*roomCountByHotel; roomIndex++ {
		for _, s := range sizes {
			roomID := rooms[roomIndex].ID
			imageID := images[index].ID
			i := generateSeedRoomImage(s, roomID, imageID)
			if err := db.Conn.Create(&i).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			if err := db.Conn.Last(&i).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			roomImages = append(roomImages, i)
			index++
		}
	}
	return roomImages
}
