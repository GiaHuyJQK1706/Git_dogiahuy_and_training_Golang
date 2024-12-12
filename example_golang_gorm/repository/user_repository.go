package repository // Bai 3

import (
	"example_golang_gorm/entity"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Insert(user *entity.User) (*entity.User, error)
	FindByID(id int64) (*entity.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

// Sử dụng GORM để thêm user
func (r *UserRepository) Insert(user *entity.User) (*entity.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Sử dụng GORM để tìm user theo ID
func (r *UserRepository) FindByID(id int64) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Preload("Projects").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
