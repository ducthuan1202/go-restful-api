package controllers

import (
	"net/http"

	"restapi/src/services"

	"github.com/gin-gonic/gin"
)

type ProductWebhookController struct {
	services *services.RootService
}

// NewProductController is function
func NewProductWebhookController(services *services.RootService) *ProductWebhookController {
	return &ProductWebhookController{
		services: services,
	}
}

// WebhookUpdate is function
func (ctrl *ProductWebhookController) UpdateStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "WebhookUpdateStatus",
		"status":  "success",
	})
}

// OrderRefun is function
func (ctrl *ProductWebhookController) OrderRefun(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "WebhookOrderRefun",
		"status":  "success",
	})
}
