package models 

import (
	"time"	
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null;unique"`
	StripeCustomerID string `gorm:"type:varchar(255);not null"`
	NoShowCount int `gorm:"default:0"`
	CancelCount string `gorm:"default:0"`
	LastSessionAt *time.Time `gorm:"default:null"`
	Notes string `gorm:"type:text"`
	Sessions []Session `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sessions"`
}