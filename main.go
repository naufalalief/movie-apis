package main

import (
	"log"
	"rest-api-gin-jwt/config"
	"rest-api-gin-jwt/docs"
	"rest-api-gin-jwt/routes"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Movie API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "movie-apis-production.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db := config.ConnectDatabase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// router
	r := routes.SetupRouter(db)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Run()
}
