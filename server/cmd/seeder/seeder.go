package main 

import (
	"log"
	"gorm.io/gorm"
	"github.com/code-guy21/TutorLink/server/internal/pkg"
	"github.com/code-guy21/TutorLink/server/internal/app/models"
)

func main() {
	config, err := pkg.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := pkg.InitDB(config)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	seedTutor(db)
}

func seedTutor(db *gorm.DB) {
	var count int64
	db.Model(&models.Tutor{}).Count(&count)

	if count > 0 {
		log.Println("Tutor profile already exists, skipping seed.")
		return
	}

	tutor := models.Tutor{
		Name: "Alexis San Javier",
		Bio: "I Specialize in Full-Stack Web Developemtnt",
		Rate: 40.00,
		Subjects: []string{"React", "JavaScript", "Node.js", "Express", "Go", "TypeScript", "SQL", "MongoDB", "GraphQL", "jQuery", "HTML", "CSS"},
		Email: "ucfknight2017@gmail.com",
		ProfilePic: "default",
	}

	if err := db.Create(&tutor).Error; err != nil {
		log.Fatalf("Failed to seed tutor profile: %v", err)
	}

	log.Println("Tutor profile seeded successfully.")
}