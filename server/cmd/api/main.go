package main

import( 
	"log"
	"github.com/code-guy21/TutorLink/server/internal/app"
)

func main(){

	application, err := app.NewApp()

	if err != nil {
		log.Fatalf("Failed to initalize app: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
	
}