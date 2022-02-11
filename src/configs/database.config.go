package configs

import (
	"fmt"
	"restapi/src/helpers"
	"restapi/src/models"
	"strings"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Instance *gorm.DB
)

// SetupDatabaseConnection is function
func SetupDatabaseConnection() {
	params := []string{
		fmt.Sprintf("host=%s", helpers.Getenv("DB_HOST")),
		fmt.Sprintf("port=%s", helpers.Getenv("DB_PORT")),
		fmt.Sprintf("user=%s", helpers.Getenv("DB_USER")),
		fmt.Sprintf("password=%s", helpers.Getenv("DB_PASSWORD")),
		fmt.Sprintf("dbname=%s", helpers.Getenv("DB_NAME")),
		"sslmode=disable",
		"TimeZone=Asia/Ho_Chi_Minh",
	}

	var err error
	Instance, err = gorm.Open(postgres.Open(strings.Join(params, " ")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
}

// CloseDatabaseConnection is function
func CloseDatabaseConnection() {
	db, err := Instance.DB()
	if err != nil {
		panic("close connection from database failed")
	}

	db.Close()
}

// GetDatabaseConnection is function
func GetDatabaseConnection() *gorm.DB {
	if Instance == nil {
		SetupDatabaseConnection()
	}
	return Instance
}

// Migration is function
func Migration() {

	Instance.AutoMigrate(&models.User{})

	Instance.AutoMigrate(&models.Category{})

	Instance.AutoMigrate(&models.Product{})

}

func Seeding() {

	// categories
	{
		cateogries := []models.Category{
			{ID: 1, Name: "Dien Thoai"},
			{ID: 2, Name: "May Tinh Bang"},
		}

		for index := range cateogries {
			Instance.Create(&cateogries[index])
		}
	}

	// products
	{
		products := []models.Product{
			{ID: 1, Name: "Iphone 11 Pro Max 128Gb", CategoryID: 1},
			{ID: 2, Name: "Samsung Note 20 Ultra", CategoryID: 1},
			{ID: 3, Name: "Xiaomi Mi 11 Lte", CategoryID: 1},
			{ID: 4, Name: "Ipad Air Wifi", CategoryID: 2},
			{ID: 5, Name: "Lenovo pad2", CategoryID: 2},
		}

		for index := range products {
			Instance.Create(&products[index])
		}
	}
}
