package repository

import (
	"database/sql"
	"example_golang_gin/entity"

	"fmt"
)

type UserRepositoryInterface interface {
	Insert(user *entity.User) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	// DI: Truyen `*sql.DB` (phụ thuộc) tu ben ngoai UserRepository.
	return &UserRepository{DB: db}
}

// Ham them du lieu user
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

// Ham tim user co id do
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE id = ?"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
