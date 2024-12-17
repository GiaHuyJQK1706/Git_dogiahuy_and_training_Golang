package repository

import (
	"project_demo/entities"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *entities.Project) error
}

type projectRepository struct {
	DB *gorm.DB
}

func (r *projectRepository) Create(project *entities.Project) error {
	return r.DB.Create(project).Error
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{DB: db}
}
