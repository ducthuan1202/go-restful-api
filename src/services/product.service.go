package services

import (
	"context"
	"errors"

	"restapi/src/dtos"
	"restapi/src/models"
	"restapi/src/repositories"

	"gorm.io/gorm"
)

type ProductServivce struct {
	Repository *repositories.ProductRepository
}

// NewProductService is function
func NewProductService(dbConnection *gorm.DB) *ProductServivce {
	return &ProductServivce{
		Repository: repositories.NewProductRepository(dbConnection),
	}
}

// GetOne is function
func (s *ProductServivce) GetOne(ctx context.Context, id int) (*models.Product, error) {

	if id == 5 {
		return s.Repository.FindById(id), nil
	}

	return nil, errors.New("product not found")
}

// Create is function
func (s *ProductServivce) Create(ctx context.Context, input dtos.ProductCreateInput) (*models.Product, error) {

	return s.Repository.CreateNew(input), nil
}

// GetAll is function
func (s *ProductServivce) GetAll(ctx context.Context, input dtos.ProductListDto) ([]*models.Product, error) {

	data := s.Repository.FindAll(input)
	return data, nil
}
