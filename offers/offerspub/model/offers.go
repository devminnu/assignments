package model

//Offers Hotel Offers
type Offers struct {
	Offers []Offer `json:"offers"`
}

//Offer Hotel Offer
type Offer struct {
	CmOfferID    string       `json:"cm_offer_id" gorm:"primary_key"`
	HotelID      string       `json:"hotel_id"`
	Hotel        Hotel        `json:"hotel" sql:"-"`
	Room         Room         `json:"room" sql:"-"`
	RatePlan     RatePlan     `json:"rate_plan" sql:"-"`
	OriginalData OriginalData `json:"original_data" gorm:"embedded"`
	Capacity     Capacity     `json:"capacity" gorm:"embedded"`
	Number       int          `json:"number"`
	Price        int          `json:"price"`
	Currency     string       `json:"currency"`
	CheckIn      string       `json:"check_in"`
	CheckOut     string       `json:"check_out"`
	Fees         []Fees       `json:"fees" gorm:"foreignKey:CmOfferID;references:CmOfferID"`
}

type OriginalData struct {
	GuaranteePolicy GuaranteePolicy `json:"GuaranteePolicy" gorm:"embedded"`
}

type GuaranteePolicy struct {
	Required bool `json:"Required" gorm:"embedded"`
}

type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}

// Maybe a Format function for printing your date
// func (j JsonBirthDate) Format(s string) string {
// 	t := time.Time(j)
// 	return t.Format(s)
// }

// type Amenity string

// func (data *Amenity) Value() (driver.Value, error) {
// 	// return data.ConvertJSONToString(), nil
// 	// if len(data) == 0 {
// 	// 	return nil, nil
// 	// }
// 	return json.RawMessage([]byte(*data)).MarshalJSON()
// }
// func (data *Amenity) Scan(value interface{}) error {
// 	// *data = data.ConvertStringToJson(valueString)
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}

// 	result := json.RawMessage{}
// 	err := json.Unmarshal(bytes, &result)
// 	*data = Amenity(result)
// 	return err

// }
