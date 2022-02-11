package controllers

import "restapi/src/services"

type RootController struct {
	Services *services.RootService

	*CategoryController
	*ProductController
	*ProductWebhookController
	*UserController
}

func NewRootController(rootServices *services.RootService) *RootController {
	return &RootController{
		Services:                 rootServices,
		ProductController:        NewProductController(rootServices),
		CategoryController:       NewCategoryController(rootServices),
		ProductWebhookController: NewProductWebhookController(rootServices),
		UserController:           NewUserController(rootServices),
	}
}
