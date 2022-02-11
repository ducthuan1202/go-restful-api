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

		// user endpoints
		usersRoutes := routes.Group("/users")
		{
			usersRoutes.POST("/register", ctrl.UserController.Create)
			usersRoutes.POST("/login", ctrl.UserController.Login)
		}

		// products endpoints
		productsRoutes := routes.Group("/products", middlewares.AuthorizeJWT(*ctrl.Services.JWTServivce))
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
