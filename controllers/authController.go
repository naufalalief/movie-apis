package controllers

import (
	"fmt"
	"net/http"
	"rest-api-gin-jwt/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Untuk login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Untuk register
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login User godoc
// @Summary Login as an user
// @Description Logging in to get jwt token for authorization
// @Tags Auth
// @Param Body body LoginInput true "body to login"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr := models.User{}
	// get user by username
	usr.Username = input.Username
	usr.Password = input.Password

	token, err := models.LoginCheck(usr.Username, usr.Password, db)
	if err != nil {
		fmt.Println("Error while login: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": usr.Username,
		"email":    usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user, "token": token})
}

// Register User godoc
// @Summary Register as an user
// @Description Registering to get jwt token for authorization
// @Tags Auth
// @Param Body body RegisterInput true "body to register"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr := models.User{}
	// get user by username
	usr.Username = input.Username
	usr.Password = input.Password
	usr.Email = input.Email

	_, err := usr.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": usr.Username,
		"email":    usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user})
}
