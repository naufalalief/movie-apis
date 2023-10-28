package models

import (
	"html"
	"rest-api-gin-jwt/utils/tokens"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Verifikasi password
func VerifyPassword(password, hashedPassword string) error {

	// Mengembalikan nilai error jika password tidak sesuai
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Check Login
func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

	// Membuat variabel error untuk menampung nilai error
	var err error

	// Membuat variabel user untuk menampung data user
	usr := User{}

	// Mengambil data user berdasarkan username
	err = db.Model(User{}).Where("username = ?", username).Take(&usr).Error
	// Mengembalikan nilai error jika terjadi kesalahan saat mengambil data user
	if err != nil {
		return "", err
	}

	// Mengembalikan nilai error jika password tidak sesuai
	err = VerifyPassword(password, usr.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Membuat token JWT
	token, err := tokens.GenerateToken(usr.ID)

	// Mengembalikan nilai error jika terjadi kesalahan saat membuat token JWT
	if err != nil {
		return "", err
	}

	// Mengembalikan nilai token
	return token, nil
}

// Save User
func (usr *User) SaveUser(db *gorm.DB) (*User, error) {

	// Membuat variabel error untuk menampung nilai error saat membuat hash password
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	// Mengembalikan nilai error jika terjadi kesalahan saat membuat hash password
	if errPassword != nil {
		return &User{}, errPassword
	}

	// Mengubah nilai password dengan nilai hashedPassword
	usr.Password = string(hashedPassword)

	// Mengubah nilai username dengan nilai yang sudah di trim spasi dan di escape html
	usr.Username = html.EscapeString(strings.TrimSpace(usr.Username))

	// Mengembalikan nilai error jika terjadi kesalahan saat membuat user
	var err error = db.Create(&usr).Error

	// Mengembalikan nilai error jika terjadi kesalahan saat membuat user
	if err != nil {
		return &User{}, err
	}

	// Mengembalikan nilai user jika tidak terjadi kesalahan
	return usr, nil
}
