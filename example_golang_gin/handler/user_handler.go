package handler // Bai 2

import (
	"example_golang_gin/usecase"
	"example_golang_gin/entity"
	"example_golang_gin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) CreateUser(c *gin.Context) { //Tao ham handle create user
	var req dto.CreateUserRequest

	// Bind JSON request vao DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mapping tu DTO sang Entity
	user := &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	createdUser, err := h.UserUsecase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mapping ta Entity sang DTO Response
	resp := dto.CreateUserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}

	c.JSON(http.StatusOK, resp)
}

// Bai 5: viet them o endpoint user/:id
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id") //id trong Param la mot phan cua duong dan

	// Validate ID la so
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}

	// Tim kiem user
	user, err := h.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

