package models

import (
	"time"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Payment struct {
	gorm.Model
	SessionID uint `gorm:"not null;uniqueIndex;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Session *Session `gorm:"foreignKey:SessionID"`
	Amount int64 `gorm:"not null"`
	Status PaymentStatus `gorm:"type:payment_status;not null;default:'pending'"`
	StripePaymentID string `gorm:"type:varchar(255);uniqueIndex"`
	RefundID *string `gorm:"type:varchar(255)"`
	RefundAmount *int64
	RefundedAt *time.Time
	PaymentMethod string `gorm:"type:varchar(50)"`
	FailureReason *string `gorm:"type:text"`
}