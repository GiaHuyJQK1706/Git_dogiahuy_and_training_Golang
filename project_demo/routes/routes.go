package routes

import (
	"project_demo/handler"

	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(router *gin.Engine, projectHandler handler.ProjectHandler) {
	projectRoutes := router.Group("/api/projects") //127.0.0.1:<port>/api/project/
	{
		projectRoutes.POST("/create", projectHandler.CreateProject)
		projectRoutes.PUT("/:id", projectHandler.UpdateProject)
		projectRoutes.DELETE("/:id", projectHandler.DeleteProject)
		projectRoutes.GET("/:id", projectHandler.GetProjectByID)
		projectRoutes.GET("/", projectHandler.ListProjects)
	}
}
