package controllers

import (
	"net/http"
	"rest-api-gin-jwt/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// inputan
type MovieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID int    `json:"age_rating_category_id"`
}

// Get All Movies godoc
// @Summary List all movies
// @Description get list of movies
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var movies []models.Movie

	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Create New Movie godoc
// @Summary Create movies
// @Description Create new movies
// @Tags Movie
// @Param Body body MovieInput true "body to create new movies"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var input MovieInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create age rating category
	movie := models.Movie{
		Title:               input.Title,
		Year:                input.Year,
		AgeRatingCategoryID: input.AgeRatingCategoryID,
	}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// save to db
	db.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})

}

// Get Movie by ID godoc
// @Summary get a movie by ID
// @Description get list of movie by ID
// @Tags Movie
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieByID(c *gin.Context) {
	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Update Movie godoc
// @Summary Update a movie by ID
// @Description Update a movie by ID
// @Tags Movie
// @Produce json
// @Param id path string true "Movie ID"
// @Param Body body MovieInput true "body to update a movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get model if exist
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// validate input
	var input MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update model
	var updatedInputMovie models.Movie

	updatedInputMovie.Title = input.Title
	updatedInputMovie.Year = input.Year
	updatedInputMovie.AgeRatingCategoryID = input.AgeRatingCategoryID
	updatedInputMovie.UpdatedAt = time.Now()

	// save to db
	db.Model(&movie).Updates(updatedInputMovie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Delete Movie godoc
// @Summary Delete a movie by ID
// @Description Delete a movie by ID
// @Tags Movie
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func DeleteMovieByID(c *gin.Context) {
	var movie models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// delete model
	db.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
