package models

import (
	"time"
	"gorm.io/gorm"
)

type SessionStatus string

const (
	StatusScheduled SessionStatus = "scheduled"
	StatusCancelled SessionStatus = "cancelled"
	StatusCompleted SessionStatus = "completed"
	StatusNoShow SessionStatus = "no_show"
)

type Session struct {
	gorm.Model
	StudentID  uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  
	Student    Student   `gorm:"foreignKey:StudentID"`  
	TutorID uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tutor Tutor `gorm:"foreignKey:TutorID"`
	BookingTime time.Time `gorm:"not null"`
	EndTime time.Time `gorm:"not null"`
	ZoomLink string `gorm:"type:varchar(255)"`
	Status SessionStatus `gorm:"type:session_status;not null;default:'scheduled'"`
	CalendlyEventID string `gorm:"type:varchar(255);uniqueIndex"`

	CancelledAt         *time.Time           `gorm:"default:null"`
    RefundAmount        *int64             
    StudentJoinedAt     *time.Time           `gorm:"default:null"`
    ActualEndTime       *time.Time           `gorm:"default:null"`
    PreviousBookingTime *time.Time           `gorm:"default:null"`
    CancellationReason  *string               `gorm:"type:text;default:''"`
    AppliedPolicyID     *uint                `gorm:"default:null;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    AppliedPolicy       CancellationPolicy   `gorm:"foreignKey:AppliedPolicyID"`
    DiscountAmount      *int64               `gorm:"default:null"`
    DiscountReason      *string               `gorm:"type:text;default:''"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
    if s.Status == "" {
        s.Status = StatusScheduled
    }

    if s.EndTime.IsZero() {
        s.EndTime = s.BookingTime.Add(time.Hour)
    }

    s.BookingTime = s.BookingTime.UTC()
    s.EndTime = s.EndTime.UTC()
    return nil
}