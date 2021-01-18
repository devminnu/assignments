package model

//Room Hotel Room
type Room struct {
	RoomID      string   `json:"room_id" gorm:"primary_key"`
	HotelID     string   `json:"hotel_id"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity" gorm:"embedded"`
}
