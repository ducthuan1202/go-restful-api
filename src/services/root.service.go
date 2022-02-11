package services

import "gorm.io/gorm"

type RootService struct {
	*ProductServivce
	*CategoryServivce
	*UserServivce
	*JWTServivce
}

func NewRootService(dbConnection *gorm.DB) *RootService {
	return &RootService{
		ProductServivce:  NewProductService(dbConnection),
		CategoryServivce: NewCategoryService(dbConnection),
		UserServivce:     NewUserService(dbConnection),
		JWTServivce:      NewJWTService(),
	}
}
