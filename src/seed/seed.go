package main

import (
	"app/src/db"
	"strings"

	"github.com/twinj/uuid"
)

func main() {
	db := db.New()
	images := seedImages(db)
	hotels := seedHotels(db)
	seedVacancies(db, hotels)
	rooms := seedRooms(db, hotels)
	seedRoomImages(db, rooms, images)
	thumbs := seedThumbs(db, hotels)
	seedThumbImages(db, thumbs, images)
}

func UuidV4() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
