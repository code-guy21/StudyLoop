package main

import( 
	"log"
	"github.com/code-guy21/TutorLink/server/internal/app"
	"github.com/code-guy21/TutorLink/server/internal/pkg"
	"github.com/code-guy21/TutorLink/server/internal/app/models"
)

func main(){
	config, err := pkg.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := pkg.InitDB(config)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&models.Tutor{})

	if err := app.Run(config); err != nil {
		log.Fatalf("Failed to run the application: %v", err)
	}
	
}