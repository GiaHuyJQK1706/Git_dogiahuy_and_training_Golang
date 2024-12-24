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
		userRoutes.GET("/:user_id/projects", projectHandler.GetProjectsByUser) // Lấy danh sách projects của user, in ra màn hình
		userRoutes.POST("/creare", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/:user_id/projects/export_csv", projectHandler.ExportProjectsByUserCSV)
	}
}

func SetupAuthRoutes(router *gin.Engine, authHandler handler.AuthHandler) {
	authRoutes := router.Group("/api/auth") // 127.0.0.1:<port>/api/auth/
	{
		authRoutes.POST("/token", authHandler.GenerateToken)
	}
}