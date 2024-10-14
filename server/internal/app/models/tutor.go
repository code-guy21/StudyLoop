package models

import (
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Name string `json:"name"` 
	Bio string `json:"bio"` 
	Email string `json:"email" gorm:"unique"`
	Rate float64 `json:"rate"`
}