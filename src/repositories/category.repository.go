package repositories

import (
	"restapi/src/dtos"
	"restapi/src/models"
	"strings"

	"gorm.io/gorm"
)

// CategoryRepository is struct
type CategoryRepository struct {
	DBConnection *gorm.DB
}

func NewCategoryRepository(dbConnection *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DBConnection: dbConnection,
	}
}

// FindAll is function
func (r *CategoryRepository) FindAll() []*models.Category {

	categories := []*models.Category{}

	r.DBConnection.
		Model(&models.Category{}).
		Preload("Products").
		Find(&categories)

	return categories
}

// CreateNew is function
func (r *CategoryRepository) CreateNew(input dtos.CategoryCreateInput) (*models.Category, error) {
	category := models.Category{
		Name: input.Name,
	}

	result := r.DBConnection.Create(&category)
	if result.RowsAffected == 1 && result.Error == nil {
		return &category, nil
	}

	return nil, result.Error
}

// DeleteById is function
func (r *CategoryRepository) DeleteById(id string) (bool, error) {
	result := r.DBConnection.Delete(&models.Category{}, id)
	if result.RowsAffected == 1 && result.Error == nil {
		return true, nil
	}
	return false, result.Error
}

// DeleteById is function
func (r *CategoryRepository) DeleteByListIds(ids string) (bool, error) {

	result := r.DBConnection.
		Delete(&models.Category{}, strings.Split(ids, ","))

	if result.RowsAffected > 0 && result.Error == nil {
		return true, nil
	}
	return false, result.Error
}
