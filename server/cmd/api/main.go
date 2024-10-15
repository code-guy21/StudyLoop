package main

import( 
	"log"
	"github.com/code-guy21/TutorLink/server/internal/app"
	"github.com/code-guy21/TutorLink/server/internal/pkg"
	"github.com/code-guy21/TutorLink/server/internal/app/models"
)

func main(){
	config, _ := pkg.LoadConfig()
	

	db, err := pkg.InitDB(config)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	
	log.Println("Database connection established")

	db.AutoMigrate(&models.Tutor{}, &models.Student{}, &models.Session{}, &models.Payment{})

	if err := app.Run(config,db); err != nil {
		log.Fatalf("Failed to run the application: %v", err)
	}
	
}