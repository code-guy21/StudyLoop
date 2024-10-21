package models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"` 
	Bio string `json:"bio"` 
	Subjects []string `gorm:"type:text[]" json:"subjects"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Rate float64 `json:"rate" gorm:"not null"`
	ProfilePic string `json:"profile_pic"`
	Sessions []Session `gorm:"foreignKey:TutorID" json:"sessions"`
}