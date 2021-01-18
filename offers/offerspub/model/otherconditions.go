package model

import (
	"strings"

	"github.com/jinzhu/gorm"
)

//OtherConditions Other terms and conditions for hotel booking and cancellation
type OtherCondition struct {
	gorm.Model
	RatePlanID string `json:"rate_plan_id" sql:"unique_index:idx_rate_plan_id_condition"`
	Name       string `json:"name" sql:"unique_index:idx_rate_plan_id_condition"`
}

//UnmarshalJSON Unmarshal OtherConditions
func (j *OtherCondition) UnmarshalJSON(b []byte) error {
	j.Name = strings.Trim(string(b), "\"")
	return nil
}
