package models

import "gorm.io/gorm"

type PolicyType string

const (
	PolicyTypeNoShow PolicyType = "no_show"
	PolicyTypeCancellation PolicyType = "cancellation"
	PolicyTypeReschedule PolicyType = "reschedule"
)

type CancellationPolicy struct {
	gorm.Model
	Type PolicyType `gorm:"type:varchar(50);not null"`
	TimeWindowMinutes int `gorm:"not null"`
	ChargePercentage int `gorm:"not null"`
	Description string `gorm:"type:text;not null"`
	IsActive bool `gorm:"default:true"`
}