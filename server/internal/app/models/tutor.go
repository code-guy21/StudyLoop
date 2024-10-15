package models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Name string `json:"name"` 
	Bio string `json:"bio"` 
	Subjects []string `gorm:"type:text[]" json:"subjects"`
	Email string `json:"email" gorm:"unique"`
	Rate float64 `json:"rate"`
	ProfilePic string `json:"profile_pic"`
}