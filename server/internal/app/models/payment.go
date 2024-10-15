package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SessionID uint `json:"sesson_id"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	StripeID string `json:"stripe_id"`
}