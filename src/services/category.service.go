package services

import (
	"context"
	"errors"

	"restapi/src/dtos"
	"restapi/src/models"
	"restapi/src/repositories"

	"gorm.io/gorm"
)

type CategoryServivce struct {
	Repository *repositories.CategoryRepository
}

// NewCategoryService is function
func NewCategoryService(dbConnection *gorm.DB) *CategoryServivce {
	return &CategoryServivce{
		Repository: repositories.NewCategoryRepository(dbConnection),
	}
}

// GetAll is function
func (s *CategoryServivce) GetAll(ctx context.Context, input dtos.CategoryListDto) ([]*models.Category, error) {

	if input.Limit <= 0 {
		return nil, errors.New("param invallid")
	}

	if input.Limit < 10 {
		data := s.Repository.FindAll()
		return data, nil
	}

	return []*models.Category{}, nil
}

// Create is function
func (s *CategoryServivce) Create(ctx context.Context, input dtos.CategoryCreateInput) (*models.Category, error) {
	return s.Repository.CreateNew(input)
}

// DeleteById is function
func (s *CategoryServivce) DeleteById(ctx context.Context, id string) (bool, error) {
	return s.Repository.DeleteById(id)
}

// DeleteMulti is function
func (s *CategoryServivce) DeleteMulti(ctx context.Context, ids string) (bool, error) {
	return s.Repository.DeleteByListIds(ids)
}
