package controllers

import "restapi/src/services"

type RootController struct {
	*CategoryController
	*ProductController
	*ProductWebhookController
}

func NewRootController(rootServices *services.RootService) *RootController {
	return &RootController{
		ProductController:        NewProductController(rootServices),
		CategoryController:       NewCategoryController(rootServices),
		ProductWebhookController: NewProductWebhookController(rootServices),
	}
}
