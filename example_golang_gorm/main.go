package main

import (
	"example_golang_gorm/entity"
	"example_golang_gorm/handler"
	"example_golang_gorm/repository"
	"example_golang_gorm/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Bai 2: Using gorm auto migrations de tao 2 bang o bai 1
	//Ket noi voi co so du lieu
	dsn := "dogiahuy:17062003@tcp(localhost:3306)/sampleDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	// Auto Migrate cho User và Project
	err = db.AutoMigrate(&entity.User{}, &entity.Project{})
	if err != nil {
		fmt.Printf("Failed to auto migrate: %v\n", err)
	}

	fmt.Println("Database migrated successfully!")

	// Khoi tao repository, usecase, handler
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Dung gin de tao server
	router := gin.Default()

	

	// Tao group endpoint
	userGroup := router.Group("/user")
	userGroup.Use(func(c *gin.Context) {
		fmt.Println("Call to endpoints user")
		c.Next()
	})

	// Tao 2 endpoint nhu yeu cau o bai 4 va bai 5
	userGroup.POST("/user_projects/create", userHandler.CreateUserWithProject) //Truy cap: 127.0.0.1:8080/user/user_projects/create
	userGroup.GET("/user_projects/:user_id", userHandler.GetUserWithProjects) //Truy cap: 127.0.0.1:8080/user/user_projects/:user_id

	// Lang nghe va chay tai cong 8080
	router.Run(":8080")
}
