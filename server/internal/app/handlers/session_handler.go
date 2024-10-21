package handlers

import (
	"net/http"
	"github.com/code-guy21/TutorLink/server/internal/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionHandler struct {
	db *gorm.DB
}

func NewSessionHandler(db *gorm.DB) *SessionHandler{
	return &SessionHandler{db: db}
}

func (h *SessionHandler) GetSessions(c *gin.Context) {
	var sessions []models.Session
	
	if err := h.db.Preload("Tutor").Preload("Student").Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sessions"})
		
	}

	c.JSON(http.StatusOK,sessions)
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	var session models.Session

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}

	if err := h.db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK,session)
}

func (h *SessionHandler) CancelSession(c *gin.Context) {
	sessionID := c.Param("id")
	var session models.Session

	if err := h.db.First(&session, sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	session.Status = "canceled"

	if err := h.db.Save(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel session"})
		return
	}

	c.JSON(http.StatusOK,session)
}