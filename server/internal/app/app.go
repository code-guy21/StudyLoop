package app 

import (
	"github.com/code-guy21/TutorLink/server/internal/app/routes"
	"github.com/code-guy21/TutorLink/server/internal/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(config *pkg.Config, db *gorm.DB) error {

	router := gin.Default()

	routes.InitRoutes(router,db)

	return router.Run(":" + config.ServerPort)
}