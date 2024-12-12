package handler

import (
	"example_golang_gorm/usecase"
	"example_golang_gorm/entity"
	"example_golang_gorm/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

// Bai 4: Hàm tạo user và project của user đó (Chỉ 1 user mới và 1 project mới)
func (h *UserHandler) CreateUserWithProject(c *gin.Context) {
	var req struct {
		User    dto.CreateUserRequest `json:"user"`
		Project struct {
			Name             string `json:"name" binding:"required"`
			ProjectStartedAt string `json:"project_started_at" binding:"required"`
		} `json:"project"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectStartedAt, err := time.Parse(time.RFC3339, req.Project.ProjectStartedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	user := &entity.User{
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Email:     req.User.Email,
		Password:  req.User.Password,
		Projects: []entity.Project{
			{
				Name:             req.Project.Name,
				ProjectStartedAt: projectStartedAt,
			},
		},
	}

	createdUser, err := h.UserUsecase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

// Bai 5: Hàm lấy thông tin user và project liên quan
func (h *UserHandler) GetUserWithProjects(c *gin.Context) {
	idParam := c.Param("user_id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	user, err := h.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
