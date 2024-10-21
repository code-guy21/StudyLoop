package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SessionID uint `json:"sesson_id" gorm:"not null"`
	Session Session `gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`  
	Amount float64 `json:"amount" gorm:"not null"`
	Currency string `json:"currency" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
	StripeID string `json:"stripe_id" gorm:"not null"`
}