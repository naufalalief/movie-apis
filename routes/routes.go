package routes

import (
	"rest-api-gin-jwt/controllers"
	"rest-api-gin-jwt/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// default gin router
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// middleware for rating
	ratingMiddlewareRoutes := r.Group("/age-rating-categories")
	ratingMiddlewareRoutes.Use(middlewares.JwtAuthMiddleware())

	ratingMiddlewareRoutes.POST("", controllers.CreateRating)
	ratingMiddlewareRoutes.PATCH("/:id", controllers.UpdateRating)
	ratingMiddlewareRoutes.DELETE("/:id", controllers.DeleteRatingByID)

	r.GET("/age-rating-categories", controllers.GetAllRating)
	r.GET("/age-rating-categories/:id", controllers.GetRatingByID)
	r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingCategoryID)

	// middleware for movies
	movieMiddlewareRoutes := r.Group("/movies")
	movieMiddlewareRoutes.Use(middlewares.JwtAuthMiddleware())

	movieMiddlewareRoutes.POST("", controllers.CreateMovie)
	movieMiddlewareRoutes.PATCH("/:id", controllers.UpdateMovie)
	movieMiddlewareRoutes.DELETE("/:id", controllers.DeleteMovieByID)

	r.GET("/movies", controllers.GetAllMovie)
	r.GET("/movies/:id", controllers.GetMovieByID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
