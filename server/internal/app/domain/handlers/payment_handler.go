 package handlers

// import (
// 	"net/http"
// 	"github.com/code-guy21/TutorLink/server/internal/app/models"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type PaymentHandler struct {
// 	db *gorm.DB
// }

// func NewPaymentHandler(db *gorm.DB) *PaymentHandler {
// 	return &PaymentHandler{db: db}
// }

// func (h *PaymentHandler) CreatePayment(c *gin.Context) {
// 	var payment models.Payment

// 	if err := c.ShouldBindJSON(&payment); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.db.Create(&payment).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
// 		return
// 	}

// 	c.JSON(http.StatusOK,payment)
// }

// func (h *PaymentHandler) HandleWebhook(c *gin.Context){
// 	c.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
// }