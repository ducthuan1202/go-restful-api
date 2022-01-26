package main

import (
	"fmt"
	"log"

	"restapi/src/configs"
	"restapi/src/controllers"
	"restapi/src/helpers"
	"restapi/src/jobs"
	"restapi/src/routers"
	"restapi/src/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func init() {

	gin.SetMode(gin.DebugMode)
	configs.SetupDatabaseConnection()
	// configs.Migration()
}

// main -> route -> controller -> service -> repository -> database
func main() {

	jobs.InitialJobs()

	// get environments
	appPort := helpers.GetenvWithDefaultValue("APP_PORT", "8080")

	rootServices := services.NewRootService(configs.GetDatabaseConnection())
	rootController := controllers.NewRootController(rootServices)

	r := gin.Default()
	r.Use(cors.Default())

	// kiem tra trang thai cua server
	r.GET("/healthz", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "OK")
	})

	// middleware log request
	// r.Use(gin.LoggerWithFormatter(middlewares.WriteLogRequestDetail))

	routers.RegisterApiRoutes(r, rootController)
	routers.RegisterWebhookRoutes(r, rootController)

	if err := r.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		log.Fatal(err.Error())
	} else {

		log.Printf("service starting at http://localhost:%s", appPort)
	}
}
