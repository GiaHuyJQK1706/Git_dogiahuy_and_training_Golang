package config

import (
	"fmt"
	"os"
	"project_demo/entities" // Import cac entity de automigrade

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDatabase - Load variables from .env and connect to the database
func InitDB() *gorm.DB {
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

	// Initialize GORM
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Perform AutoMigrate for all entities
	err = DB.AutoMigrate(
		&entities.Project{}, // Add your entities here
		// Add more entities as needed
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	fmt.Println("🚀 Database connected successfully! :=D")

	return DB
}
