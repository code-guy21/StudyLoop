package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SessionID uint
	Amount int64 `gorm:"not null"`
	Status string `gorm:"not null"`
	StripeID string `gorm:"not null"`
}