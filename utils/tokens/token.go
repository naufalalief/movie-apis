package tokens

import (
	"fmt"
	"rest-api-gin-jwt/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// API_SECRET adalah secret key untuk JWT
var API_SECRET = utils.GetEnv("API_SECRET", "secret")

// GenerateToken adalah fungsi untuk membuat token JWT
func GenerateToken(user_id uint) (string, error) {
	// Mengambil nilai dari environment variable
	token_hour_lifespan, err := strconv.Atoi(utils.GetEnv("TOKEN_HOUR_LIFESPAN", "1"))
	if err != nil {
		return "", err
	}
	// Membuat payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_hour_lifespan)).Unix()
	// Membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Membuat string token
	return token.SignedString([]byte(API_SECRET))
}

// Ekstrak token JWT adalah fungsi untuk mengekstrak token JWT
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	// Mengambil nilai token dari header
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// TokenValid adalah fungsi untuk mengecek apakah token valid
func TokenValid(c *gin.Context) error {
	// Mengambil token
	tokenString := ExtractToken(c)
	// Membuat token
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Mengecek apakah token valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Mengembalikan nilai API_SECRET
		return []byte(API_SECRET), nil
	})
	// Mengembalikan error
	if err != nil {
		return err
	}
	return nil
}

// Ekstrak token ID adalah fungsi untuk mengekstrak token ID
func ExtractTokenID(c *gin.Context) (uint, error) {
	// Mengambil token
	tokenString := ExtractToken(c)
	// Membuat token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Mengecek apakah token valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Mengembalikan nilai API_SECRET
		return []byte(API_SECRET), nil
	})
	// Mengembalikan error
	if err != nil {
		return 0, err
	}
	// Mengembalikan nilai user_id
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Mengambil nilai user_id
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	// Mengembalikan nilai 0
	return 0, nil
}
