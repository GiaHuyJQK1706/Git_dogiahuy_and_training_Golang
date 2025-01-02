package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET")) // Lấy secret từ biến môi trường

// GenerateJWT - Tao token JWT
func GenerateJWT(userID uint) (string, error) {
	// Định nghĩa các claims của token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token het sau 24 gio
	}

	// Tao token voi claims va signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ky token voi secret key
	return token.SignedString(jwtSecretKey)
}

// ValidateJWT - Xac thuc token JWT
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecretKey, nil
	})
}

// HashPassword - Ma hoa mat khau bang cach bam
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

// CheckPasswordHash - Kiem tra mat khau co trung khop voi hash hay khong
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	return err == nil
}
