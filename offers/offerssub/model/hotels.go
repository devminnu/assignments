package model

//Hotel Hotel
type Hotel struct {
	HotelID     string    `json:"hotel_id" gorm:"primary_key"`
	Room        []Room    `json:"room" gorm:"foreignKey:HotelID;references:HotelID"`
	RatePlan    RatePlan  `json:"rate_plan" gorm:"foreignKey:HotelID;references:HotelID"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Telephone   string    `json:"telephone"`
	Amenities   []Amenity `json:"amenities" gorm:"foreignKey:HotelID;references:HotelID"`
	Description string    `json:"description"`
	RoomCount   int       `json:"room_count"`
	Currency    string    `json:"currency"`
	Fees        []Fee     `json:"fees" gorm:"foreignKey:HotelID;references:HotelID"`
}
