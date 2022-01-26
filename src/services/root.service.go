package services

import "gorm.io/gorm"

type RootService struct {
	ProductServivce  *ProductServivce
	CategoryServivce *CategoryServivce
}

func NewRootService(dbConnection *gorm.DB) *RootService {
	return &RootService{
		ProductServivce:  NewProductService(dbConnection),
		CategoryServivce: NewCategoryService(dbConnection),
	}
}
