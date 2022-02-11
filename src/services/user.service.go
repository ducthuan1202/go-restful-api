package services

import (
	"context"

	"restapi/src/dtos"
	"restapi/src/models"
	"restapi/src/repositories"

	"gorm.io/gorm"
)

type UserServivce struct {
	Repository *repositories.UserRepository
}

// NewCategoryService is function
func NewUserService(dbConnection *gorm.DB) *UserServivce {
	return &UserServivce{
		Repository: repositories.NewUserRepository(dbConnection),
	}
}

// Create is function
func (s *UserServivce) Create(ctx context.Context, input dtos.UserCreateInput) (*models.User, error) {
	return s.Repository.CreateUser(input)
}
