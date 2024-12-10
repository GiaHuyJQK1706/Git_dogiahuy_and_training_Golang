package main

import (
	"database/sql"
	"example_golang_gin/handler"
	"example_golang_gin/repository"
	"example_golang_gin/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	//Ket noi voi co so du lieu
	db, err := sql.Open("mysql", "dogiahuy:17062003@tcp(localhost:3306)/sampleDB?parseTime=true")
	if err != nil {
		panic(fmt.Errorf("connection fail: %v", err))
	}
	defer db.Close()

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`
	
	_, err = db.Exec(query)
	if err != nil {
		panic(fmt.Errorf("create table fail: %v", err))
	}

	fmt.Println("Connection success, create table success")

	// Khoi tao repository, usecase, handler (Bai 2)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Dung gin de tao server (Bai 1) 
	router := gin.Default()

	// Tao endpoint POST /user/create/ (Bai 2)
	/*
		router.POST("/user/create", userHandler.CreateUser)
	*/

	// Bai 6:
	userGroup := router.Group("/user")
	userGroup.Use(func(c *gin.Context) {
		fmt.Println("Call to endpoints user")
		c.Next()
	})

	userGroup.POST("/create", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUserByID)


	// Lang nghe va chay tai cong 8080 (Bai 1)
	router.Run(":8080")
}
