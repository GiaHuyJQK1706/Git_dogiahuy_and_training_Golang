package repository

import (
	"project_demo/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(project *entities.User) error
	Update(project *entities.User) error
	Delete(id uint) error
	GetByID(id uint) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

// Tao nguoi dung
func (r *userRepository) Create(user *entities.User) error {
	return r.DB.Create(user).Error
}

// Cap nhat nguoi dung
func (r *userRepository) Update(user *entities.User) error {
	return r.DB.Save(user).Error
}

// Xoa nguoi dung
func (r *userRepository) Delete(id uint) error {
	return r.DB.Delete(&entities.User{}, id).Error
}

// Lay (Xem chi tiet) user theo ID
func (r *userRepository) GetByID(id uint) (entities.User, error) {
	var user entities.User
	err := r.DB.First(&user, id).Error
	
	return user, err
}

// Tim user theo email
func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error

	return user, err
}
// DI
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
