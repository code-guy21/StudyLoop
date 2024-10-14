package models

import (
	"time"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	TutorID uint `gorm:"not null"`
	StudentName string `gorm:"not null"`
	ScheduledAt time.Time `gorm:"not null"`
	Status string `gorm:"not null"`
	ZoonLink string `gorm:"not null"`
}