package main

import (
	"example_golang_gorm/entity"
	"example_golang_gorm/handler"
	"example_golang_gorm/repository"
	"example_golang_gorm/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Bai 2: Using gorm auto migrations de tao 2 bang o bai 1
	/* Tao bien DSN theo kieu hardcode
	dsn := "dogiahuy:17062003@tcp(localhost:3306)/sampleDB?charset=utf8mb4&parseTime=True&loc=Local"
	*/

	// Load file .env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// Lay cac bien moi truong tu file .env
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	// Tao bien DSN cac bien moi truong
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc)

	// In ra DSN (debug)
	fmt.Println("DSN:", dsn)

	// Dung gorm ket noi voi co so du lieu
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
