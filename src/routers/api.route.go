package routers

import (
	"restapi/src/controllers"
	"restapi/src/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine, ctrl *controllers.RootController) *gin.RouterGroup {

	routes := r.Group("/api/v1")
	{
		routes.Use(middlewares.WriteLogApiRequest)

		// products endpoints
		productsRoutes := routes.Group("/products")
		{
			productsRoutes.GET("/", ctrl.ProductController.GetAll)
			productsRoutes.GET("/:productId", ctrl.ProductController.GetById)
			productsRoutes.POST("/", ctrl.ProductController.Create)
			productsRoutes.PATCH("/", ctrl.ProductController.Update)
			productsRoutes.PUT("/", ctrl.ProductController.Update)
			productsRoutes.DELETE("/:productId", ctrl.ProductController.Delete)
		}

		// categories endpoints
		categoriesRoutes := routes.Group("/categories")
		{
			categoriesRoutes.GET("/", ctrl.CategoryController.GetAll)
			categoriesRoutes.GET("/:id", ctrl.CategoryController.GetById)
			categoriesRoutes.POST("/", ctrl.CategoryController.Create)
			categoriesRoutes.PATCH("/", ctrl.CategoryController.Update)
			categoriesRoutes.PUT("/", ctrl.CategoryController.Update)
			categoriesRoutes.DELETE("/:id", ctrl.CategoryController.DeleteById)
			categoriesRoutes.DELETE("/delete", ctrl.CategoryController.Delete)
		}

	}

	return routes
}
