package tests

// test cho repository/user_repository.go

import (
	"project_demo/entities"
	"project_demo/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestUserDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&entities.User{})
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestUserDB()
	repo := repository.NewUserRepository(db)
	user := entities.User{Name: "Test User"}

	err := repo.Create(&user)

	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
}

func TestGetUserByID(t *testing.T) {
	db := setupTestUserDB()
	repo := repository.NewUserRepository(db)
	user := entities.User{Name: "Test User"}
	db.Create(&user)

	result, err := repo.GetByID(user.ID)

	assert.Nil(t, err)
	assert.Equal(t, "Test User", result.Name)
}
