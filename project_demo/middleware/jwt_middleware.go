package middleware

import (
	"net/http"
	"project_demo/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware - Middleware xac thuc JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lay token tu header Authorization
		tokenString := c.GetHeader("Authorization")

		// Kiem tra token co duoc gui kem them hay khong
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Kiem tra tien to "Bearer " (Xau "Bearer " co do dai la 7)
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Xac thuc token
		token, err := config.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Lay thong tin user_id tu token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Luu user_id vao context de dung trong handler
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()
	}
}
