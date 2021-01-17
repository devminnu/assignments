package model

import (
	"gorm.io/gorm"
)

//Room Hotel Room
type Room struct {
	// Offers      []*Offer `json:"offers" gorm:"foreignKey:RoomID;references:RoomID"`
	RoomID      string   `json:"room_id" gorm:"primary_key"`
	HotelID     string   `json:"hotel_id"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity" gorm:"embedded"`
}

type User struct {
	// gorm.Model
	UserID     string `gorm:"primary_key"`
	Name       string
	Age        int
	CreditCard []CreditCard `gorm:"foreignKey:UserID;references:UserID"`
}

type CreditCard struct {
	gorm.Model
	UserID string
	Number string
}
