package model

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// Amenity Hotel amenities
type Amenity struct {
	gorm.Model
	HotelID string `json:"hotel_id"`
	Name    string `json:"name"`
}

// Implement Marshaler and Unmarshaler interface
func (j *Amenity) UnmarshalJSON(b []byte) error {
	/*
		s := strings.Trim(string(b), "\"")
			t, err := time.Parse("2006-01-02", s)
					if err != nil {
						return err
					}
					*j = JsonBirthDate(t)
					json.Unmarshal(b, j)
	*/
	// fmt.Println("mk", strings.Trim(string(b), "\""))
	j.Name = strings.Trim(string(b), "\"")
	return nil
}

// func (j Amenity) MarshalJSON() ([]byte, error) {
// 	fmt.Println("MarshalJSON")
// 	return json.Marshal(j)
// }
