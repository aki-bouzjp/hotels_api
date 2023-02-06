package main

import (
	"app/src/db"
	"app/src/model"
	"math/rand"
	"time"

	"fmt"
)

func generateSeedVacancy(hotelId string, empty bool) *model.Vacancy {
	return &model.Vacancy{
		ID:      UuidV4(),
		HotelID: hotelId,
		Empty:   empty,
	}
}

func seedVacancies(db *db.DB, hotels []*model.Hotel) []*model.Vacancy {
	vacancies := []*model.Vacancy{}
	for hotelIndex := 0; hotelIndex < hotelTotalCount; hotelIndex++ {
		var empty = false
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) < 50 {
			empty = true
		}
		hotelID := hotels[hotelIndex].ID
		v := generateSeedVacancy(hotelID, empty)
		if err := db.Conn.Create(&v).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
		if err := db.Conn.Last(&v).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
	}
	return vacancies
}
