package models

import (
	"time"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	TutorID uint `json:"tutor_id"`
	StartTime time.Time `json:"start_time"`
	StudentID  uint      `json:"student_id"`    
    Student    Student   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EndTime time.Time `json:"end_time"`
	Status string `json:"string"`
	PaymentID string `json:"payment_id"`
	ZoomLink string `json:"zoom_link"`
}
