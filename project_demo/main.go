package main

import (
	"project_demo/config"
	"project_demo/handler"
	"project_demo/repository"
	"project_demo/routes"
	"project_demo/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db := config.InitDB()

	// Initialize repository, usecase, and handler
	// Initialize repository: projectRepo, userRepo
	projectRepo := repository.NewProjectRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize usecase: projectUC, userUC
	projectUC := usecase.NewProjectUsecase(projectRepo)
	userUC := usecase.NewUserUsecase(userRepo)

	// Initialize handler: projectHandler, userHandler
	projectHandler := handler.ProjectHandler{ProjectUC: projectUC}
	userHandler := handler.UserHandler{UserUC: userUC}

	// Setup routes
	router := gin.Default()
	routes.SetupProjectRoutes(router, projectHandler, userHandler)
	routes.SetupAuthRoutes(router, handler.AuthHandler{UserRepo: userRepo})

	// Start server
	router.Run(":8080")
}
