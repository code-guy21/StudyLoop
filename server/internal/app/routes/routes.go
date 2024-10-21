package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/code-guy21/TutorLink/server/internal/app/handlers"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	tutorHandler := handlers.NewTutorHandler(db)
	sessionHandler := handlers.NewSessionHandler(db)
	paymentHandler := handlers.NewPaymentHandler(db)

	router.GET("/tutor/:id", tutorHandler.GetTutorProfile)

	router.GET("/sessions", sessionHandler.GetSessions)
	router.POST("/sessions", sessionHandler.CreateSession)
	router.PUT("/sessions/:id/cancel", sessionHandler.CancelSession)

	router.POST("/payments", paymentHandler.CreatePayment)
	router.POST("/webhooks/stripe", paymentHandler.HandleWebhook)
}