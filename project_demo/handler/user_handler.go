package handler

import (
	"net/http"
	"project_demo/dto"
	"project_demo/usecase"
	"project_demo/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUC usecase.UserUsecase
}

// API tao user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Custom validation, check password
	if err := validator.ValidatePassword(request.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check xong validation
	
	// Cac cau lenh duoi duoc thuc hien khi validation check thanh cong (Check xong validation)
	user, err := h.UserUC.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// API cap nhat user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUC.UpdateUser(uint(id), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// API xoa user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err := h.UserUC.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// API Lay (Xem chi tiet) user theo ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	project, err := h.UserUC.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}
