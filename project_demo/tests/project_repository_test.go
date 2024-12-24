package tests

// test cho repository/project_repository.go

import (
	"project_demo/entities"
	"project_demo/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&entities.Project{})
	return db
}

func TestCreateProject(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewProjectRepository(db)
	project := entities.Project{Name: "Test Project"}

	err := repo.Create(&project)

	assert.Nil(t, err)
	assert.NotZero(t, project.ID)
}

func TestGetProjectByID(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewProjectRepository(db)
	project := entities.Project{Name: "Test Project"}
	db.Create(&project)

	result, err := repo.GetByID(project.ID)

	assert.Nil(t, err)
	assert.Equal(t, "Test Project", result.Name)
}
