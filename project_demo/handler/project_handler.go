package handler

import (
	"net/http"
	"project_demo/dto"
	"project_demo/usecase"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectUC usecase.ProjectUsecase
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var request dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.ProjectUC.CreateProject(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
