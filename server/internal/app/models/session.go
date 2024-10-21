package models

import (
	"time"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	TutorID uint `json:"tutor_id" gorm:"not null"`
	Tutor Tutor `gorm:"foreignKey:TutorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	StudentID  uint      `json:"student_id" gorm:"not null"`    
    Student    Student   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EndTime time.Time `json:"end_time" gorm:"not null"`
	Status string `json:"string" gorm:"not null"`
	PaymentID string `json:"payment_id"`
	ZoomLink string `json:"zoom_link"`
}
