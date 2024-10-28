package models

// import (
// 	"gorm.io/gorm"
// 	"github.com/lib/pq"
// )

// type Tutor struct {
// 	gorm.Model
// 	Name string `json:"name" gorm:"not null"` 
// 	Bio string `json:"bio"` 
// 	Subjects pq.StringArray `gorm:"type:text[]" json:"subjects"`
// 	Email string `json:"email" gorm:"uniqueIndex;not null"`
// 	Rate float64 `json:"rate" gorm:"not null"`
// 	ProfilePic string `json:"profile_pic"`
// 	Sessions []Session `gorm:"foreignKey:TutorID" json:"sessions"`
// }