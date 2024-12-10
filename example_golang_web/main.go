package main

import (
	"database/sql"
	"fmt"
	"time"

	"example_golang_web/usecase"
	"example_golang_web/repository"
	"example_golang_web/entity"

	_ "github.com/go-sql-driver/mysql" //Dung gian tiep goi/thu vien github.com/go-sql-driver/mysql
)

func main() {
	db, err := sql.Open("mysql", "dogiahuy:17062003@tcp(localhost:3306)/sampleDB")
	if err != nil {
		panic(fmt.Errorf("connection fail: %v", err))
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	//Bai tap 1
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

	//Bai tap 4
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	user := &entity.User{
		FirstName: "Do",
		LastName:  "Huy",
		Email:     "dogiahuy@example.com",
		Password:  "password123",
	}

	createdUser, err := userUsecase.CreateUser(user)
	if err != nil {
		panic(fmt.Errorf("create user fail: %v", err))
	}

	fmt.Printf("User created successfully: %+v\n", createdUser)
}
