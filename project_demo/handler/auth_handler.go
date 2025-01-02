package handler

import (
	"net/http"
	"project_demo/config"
	"project_demo/repository"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo repository.UserRepository
}

// RequestBody - Định nghĩa dữ liệu đầu vào
type TokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// POST /token
func (h *AuthHandler) GenerateToken(c *gin.Context) {
	var req TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tìm user theo email
	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Kiểm tra mật khẩu (giả sử dùng bcrypt)
	if !config.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Tạo JWT token
	token, err := config.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Trả về token
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
