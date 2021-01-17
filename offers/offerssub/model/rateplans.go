package model

//RatePlan Hotel's Rate Plan
type RatePlan struct {
	RatePlanID         string               `json:"rate_plan_id" gorm:"primary_key"`
	HotelID            string               `json:"hotel_id"`
	CancellationPolicy []CancellationPolicy `json:"cancellation_policy" gorm:"foreignKey:RatePlanID;references:RatePlanID"`
	Name               string               `json:"name"`
	OtherConditions    []OtherCondition     `json:"other_conditions" gorm:"foreignKey:RatePlanID;references:RatePlanID"`
	MealPlan           string               `json:"meal_plan"`
}
