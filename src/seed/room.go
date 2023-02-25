package main

import (
	"app/src/db"
	"app/src/model"
	"strconv"

	"fmt"
)

var roomCountByHotel = 9

func generateSeedRoom(name string, groupName string, hotelId string) *model.Room {
	return &model.Room{
		ID:        UuidV4(),
		Name:      name,
		GroupName: groupName,
		HotelID:   hotelId,
	}
}

func seedRooms(db *db.DB, hotels []*model.Hotel) []*model.Room {
	index := 0
	rooms := []*model.Room{}
	// total: 45, 各ホテルに9つずつルーム
	for hotelIndex := 0; hotelIndex < hotelTotalCount; hotelIndex++ {
		for roomIndex := 0; roomIndex < roomCountByHotel; roomIndex++ {
			hotelID := hotels[hotelIndex].ID
			r := generateSeedRoom("test room "+strconv.Itoa(hotelIndex+1), "", hotelID)
			if err := db.Conn.Create(&r).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			if err := db.Conn.Last(&r).Error; err != nil {
				fmt.Printf("%+v", err)
				continue
			}
			rooms = append(rooms, r)
			index++
		}
	}
	return rooms
}
