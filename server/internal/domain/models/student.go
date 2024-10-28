package models 

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Sessions []Session `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sessions"`
}