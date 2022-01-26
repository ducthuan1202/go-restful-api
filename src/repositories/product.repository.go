package repositories

import (
	"fmt"
	"restapi/src/dtos"
	"restapi/src/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DBConnection *gorm.DB
}

func NewProductRepository(dbConnection *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DBConnection: dbConnection,
	}
}

// FindById is function
func (r *ProductRepository) FindById(id int) *models.Product {
	return &models.Product{
		Name: fmt.Sprintf("Product %d", id),
	}
}

// CreateNew is function
func (r *ProductRepository) CreateNew(input dtos.ProductCreateInput) *models.Product {
	return &models.Product{
		Name: input.Name,
	}
}

// CreateNew is function
func (r *ProductRepository) FindAll(input dtos.ProductListDto) []*models.Product {

	products := []*models.Product{}

	r.DBConnection.
		Model(&models.Product{}).
		Preload("Category").
		Limit(int(input.Limit)).
		Find(&products)

	return products
}
