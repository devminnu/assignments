package model

import "github.com/jinzhu/gorm"

//CancellationPolicy Hotel room cancellation policies
type CancellationPolicy struct {
	gorm.Model
	RatePlanID        string `json:"rate_plan_id"`
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}
