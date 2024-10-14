package app 

import (
	"log"
	"github.com/code-guy21/TutorLink/server/internal/app/routes"
	"github.com/code-guy21/TutorLink/server/internal/pkg"
	"github.com/gin-gonic/gin"
)

func Run(config *pkg.Config) error {
	db, err := pkg.InitDB(config)

	if err != nil {
		return err
	}

	log.Println("Database connection established")

	router := gin.Default()

	routes.InitRoutes(router,db)

	return router.Run(":" + config.ServerPort)
}