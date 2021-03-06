package model

import "gorm.io/gorm"

//Fees Hotel offer fees
type Fee struct {
	gorm.Model
	CmOfferID   string  `json:"cm_offer_id"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float64 `json:"percent"`
}
