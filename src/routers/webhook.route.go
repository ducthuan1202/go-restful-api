package routers

import (
	"restapi/src/controllers"
	"restapi/src/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterWebhookRoutes(r *gin.Engine, ctrl *controllers.RootController) *gin.RouterGroup {
	routes := r.Group("/webhook")
	{
		routes.Use(middlewares.WriteLogWebhookRequest)

		routes.GET("/order/:orderId/update-status", ctrl.ProductWebhookController.UpdateStatus)

		routes.GET("/orders/:orderId/refun", ctrl.ProductWebhookController.OrderRefun)
	}

	return routes
}
