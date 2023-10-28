package config

import (
	"fmt"
	"rest-api-gin-jwt/models"
	"rest-api-gin-jwt/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.GetEnv("DATABASE_USERNAME", "root")
	password := utils.GetEnv("DATABASE_PASSWORD", "12356")
	host := utils.GetEnv("DATABASE_HOST", "127.0.0.1")
	port := utils.GetEnv("DATABASE_PORT", "3306")
	database := utils.GetEnv("DATABASE_NAME", "database_movies")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Movie{}, &models.AgeRatingCategory{})

	return db
}
