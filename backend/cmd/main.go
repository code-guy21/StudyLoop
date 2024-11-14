package main 

import (
	"log"
	"github.com/code-guy21/StudyLoop/backend/config"
	"github.com/code-guy21/StudyLoop/backend/internal/routes"
	"github.com/code-guy21/StudyLoop/backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func main(){
	config.LoadConfig()

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
