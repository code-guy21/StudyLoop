package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/code-guy21/TutorLink/server/internal/app/handlers"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	tutorHandler := handlers.NewTutorHandler(db)

	router.GET("/tutor", tutorHandler.GetTutorProfile)
}