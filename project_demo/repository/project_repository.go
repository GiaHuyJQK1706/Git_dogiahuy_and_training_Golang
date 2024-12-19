package repository

import (
	"project_demo/entities"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *entities.Project) error
	Update(project *entities.Project) error
	Delete(id uint) error
	GetByID(id uint) (entities.Project, error)
	List() ([]entities.Project, error)
	GetProjectsByUserID(userID uint) ([]entities.Project, error)
}

type projectRepository struct {
	DB *gorm.DB
}

// Tao project
func (r *projectRepository) Create(project *entities.Project) error {
	return r.DB.Create(project).Error
}

// Cap nhat project
func (r *projectRepository) Update(project *entities.Project) error {
	return r.DB.Save(project).Error
}

// Xoa project
func (r *projectRepository) Delete(id uint) error {
	return r.DB.Delete(&entities.Project{}, id).Error
}

// Lay (Xem chi tiet) project theo ID
func (r *projectRepository) GetByID(id uint) (entities.Project, error) {
	var project entities.Project
	err := r.DB.First(&project, id).Error
	return project, err
}

// Lay (Xem) danh sach tat ca cac project
func (r *projectRepository) List() ([]entities.Project, error) {
	var projects []entities.Project
	err := r.DB.Find(&projects).Error
	return projects, err
}

// Lay (Xem) danh sach cac project theo nguoi dung
func (r *projectRepository) GetProjectsByUserID(userID uint) ([]entities.Project, error) {
	var projects []entities.Project
	err := r.DB.Preload("Users").Where("id IN (?)",
		r.DB.Table("project_users").Select("project_id").Where("user_id = ?", userID),
	).Find(&projects).Error
	return projects, err
}

// DI
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{DB: db}
}
