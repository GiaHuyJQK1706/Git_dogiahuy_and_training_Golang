package handler

import (
	"net/http"
	"project_demo/dto"
	"project_demo/usecase"
	"project_demo/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectUC usecase.ProjectUsecase
}

// API tao project
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var request dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Custom validation, check validation
	if err := validator.ValidateCategory(request.Category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateProjectDates(request.ProjectStartedAt, *request.ProjectEndedAt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check xong validation
	
	// Cac cau lenh duoi duoc thuc hien khi validation check thanh cong (Check xong validation)
	project, err := h.ProjectUC.CreateProject(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// API cap nhat project
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var request dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.ProjectUC.UpdateProject(uint(id), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// API xoa project
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ProjectUC.DeleteProject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// API Lay (Xem chi tiet) project theo ID
func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	project, err := h.ProjectUC.GetProjectByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// API lay (Xem) danh sach tat ca cac project
func (h *ProjectHandler) ListProjects(c *gin.Context) {
	projects, err := h.ProjectUC.ListProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// API lay (Xem) danh sach cac project theo nguoi dung
func (h *ProjectHandler) GetProjectsByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	projects, err := h.ProjectUC.GetProjectsByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}
