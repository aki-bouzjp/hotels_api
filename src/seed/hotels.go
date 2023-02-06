package main

import (
	"app/src/db"
	"app/src/model"

	"fmt"
	"strconv"
)

var hotelTotalCount = 5

func generateSeedHotel(happyHotelId string, name string) *model.Hotel {
	return &model.Hotel{
		ID:              UuidV4(),
		HappyHotelID:    happyHotelId,
		HotelType:       "",
		Lng:             139.7277816,
		Lat:             35.6575055,
		Name:            name,
		Address:         "test address",
		Url:             "test url",
		Score:           1.0,
		SinglePrice:     1000.0,
		RestMinPrice:    3000.0,
		LodgingMinPrice: 5000.0,
		PricingPlans:    "{}",
		OnlineBooking:   true,
		OnlineChecking:  true,
	}
}

func seedHotels(db *db.DB) []*model.Hotel {
	hotels := []*model.Hotel{}
	for hotelIndex := 0; hotelIndex < hotelTotalCount; hotelIndex++ {
		h := generateSeedHotel(strconv.Itoa(hotelIndex+1), "test hotel name "+strconv.Itoa(hotelIndex+1))
		if err := db.Conn.Create(&h).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
		if err := db.Conn.Last(&h).Error; err != nil {
			fmt.Printf("%+v", err)
			continue
		}
		hotels = append(hotels, h)
	}
	return hotels
}
