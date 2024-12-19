package routes

import (
	"project_demo/handler"

	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(router *gin.Engine, projectHandler handler.ProjectHandler, userHandler handler.UserHandler) {
	projectRoutes := router.Group("/api/projects") // 127.0.0.1:<port>/api/project/
	{
		projectRoutes.POST("/create", projectHandler.CreateProject)
		projectRoutes.PUT("/:id", projectHandler.UpdateProject)
		projectRoutes.DELETE("/:id", projectHandler.DeleteProject)
		projectRoutes.GET("/:id", projectHandler.GetProjectByID)
		projectRoutes.GET("/", projectHandler.ListProjects)
	}

	userRoutes := router.Group("/api/users") // 127.0.0.1:<port>/api/user/
	{
		userRoutes.GET("/:user_id/projects", projectHandler.GetProjectsByUser) // Lấy danh sách projects của user
		userRoutes.POST("/creare", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		projectRoutes.DELETE("/:id", userHandler.DeleteUser)
		projectRoutes.GET("/:id", userHandler.GetUserByID)
	}
}
