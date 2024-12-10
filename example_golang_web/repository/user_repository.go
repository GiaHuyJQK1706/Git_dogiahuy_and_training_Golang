package repository

import (
	"database/sql"
	"example_golang_web/entity"
)

type UserRepositoryInterface interface {
	Insert(user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	// DI: Truyen `*sql.DB` (phụ thuộc) tu ben ngoai UserRepository. (Bai tap 3)
	return &UserRepository{DB: db}
}

func (r *UserRepository) Insert(user *entity.User) (*entity.User, error) {
	query := "INSERT INTO users (first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	result, err := r.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}
