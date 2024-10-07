package repositories

import (
	"fmt"
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/code-guy21/TutorLink/server/internal/models"
)

var DB *gorm.DB

func InitDatabse(){
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Session{}, &models.Payment{})

	DB = db

	fmt.Println("Database connection successful!")
}