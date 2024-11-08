package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"` 
	Bio string `gorm:"type:text`
	Subjects pq.StringArray `gorm:"type:text[]" json:"subjects"`
	Email string `gorm:"type:varchar(255);not null;unique"`
	Rate int64 `gorm:"not null"`
	IsAvailable bool `gorm:"default:true"`	
	CalendlyURL string `gorm:"type:varchar(255)`
	ZoomURL string `gorm:"type:varchar(255)"`
	ProfilePic string `json:"profile_pic"`
	Sessions []Session `gorm:"foreignKey:TutorID" json:"sessions"`
}