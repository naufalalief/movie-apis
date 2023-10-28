package controllers

import (
	"net/http"
	"rest-api-gin-jwt/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// inputan
type AgeRatingCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Age Rating Categories godoc
// @Summary List all age rating categories
// @Description get list of age rating categories
// @Tags Age Rating Category
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var ratings []models.AgeRatingCategory

	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// Create New Age Rating Category godoc
// @Summary Create age rating categories
// @Description Create new age rating categories
// @Tags Age Rating Category
// @Param Body body AgeRatingCategoryInput true "body to create new age rating categories"
// @Param Authorization header string true "Authorization. How to input in swagger: `Bearer <token>`"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func CreateRating(c *gin.Context) {
	var input AgeRatingCategoryInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create age rating category
	rating := models.AgeRatingCategory{
		Name:        input.Name,
		Description: input.Description,
	}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// save to db
	db.Create(&rating)
	c.JSON(http.StatusOK, gin.H{"data": rating})

}

// Get Age Rating Category by ID godoc
// @Summary get an age rating category by ID
// @Description get list of age rating category by ID
// @Tags Age Rating Category
// @Produce json
// @Param id path string true "Age Rating Category ID"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [get]
func GetRatingByID(c *gin.Context) {
	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get Movies from Age Rating Category godoc
// @Summary get movies from an age rating category by ID
// @Description get list of movies from an age rating category by ID
// @Tags Age Rating Category
// @Produce json
// @Param id path string true "Age Rating Category ID"
// @Success 200 {object} []models.Movie
// @Router /age-rating-categories/{id}/movies [get]
func GetMoviesByRatingCategoryID(c *gin.Context) {
	var movies []models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Update Rating godoc
// @Summary Update an age rating category by ID
// @Description Update an age rating category by ID
// @Tags Age Rating Category
// @Produce json
// @Param id path string true "Age Rating Category ID"
// @Param Body body AgeRatingCategoryInput true "body to update age rating categories"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [patch]
func UpdateRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get model if exist
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// validate input
	var input AgeRatingCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update model
	var updatedInputRating models.AgeRatingCategory

	updatedInputRating.Name = input.Name
	updatedInputRating.Description = input.Description
	updatedInputRating.UpdatedAt = time.Now()

	// save to db
	db.Model(&rating).Updates(updatedInputRating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Delete Rating godoc
// @Summary Delete an age rating category by ID
// @Description Delete an age rating category by ID
// @Tags Age Rating Category
// @Produce json
// @Param id path string true "Age Rating Category ID"
// @Success 200 {object} map[string]boolean
// @Router /age-rating-categories/{id} [delete]
func DeleteRatingByID(c *gin.Context) {
	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// delete model
	db.Delete(&rating)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
