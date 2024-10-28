package handlers

// import (
// 	"net/http"
// 	"github.com/code-guy21/TutorLink/server/internal/app/models"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type TutorHandler struct {
// 	db *gorm.DB
// }

// func NewTutorHandler(db *gorm.DB) *TutorHandler {
// 	return &TutorHandler{db: db}
// }

// func (h *TutorHandler) GetTutorProfile(c *gin.Context) {
// 	var tutor models.Tutor

// 	tutorID := c.Param("id")

// 	if err := h.db.First(&tutor, tutorID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Tutor not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, tutor)
// }