package usecase

import (
	"project_demo/dto"
	"project_demo/entities"
	"project_demo/repository"
)

type UserUsecase interface {
	CreateUser(dto.CreateUserRequest) (entities.User, error)
	UpdateUser(id uint, request dto.UpdateUserRequest) (entities.User, error)
	DeleteUser(id uint) error
	GetUserByID(id uint) (entities.User, error)
}

type userUsecase struct {
	UserRepo repository.UserRepository
}

// Tao project moi
func (uuc *userUsecase) CreateUser(request dto.CreateUserRequest) (entities.User, error) {
	user := entities.User{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
		CreatedAt: request.CreatedAt,
	}

	err := uuc.UserRepo.Create(&user)
	return user, err
}

// Cap nhat project
func (uuc *userUsecase) UpdateUser(id uint, request dto.UpdateUserRequest) (entities.User, error) {
	user, err := uuc.UserRepo.GetByID(id)
	if err != nil {
		return user, err
	}

	user.Name = request.Name
	user.Email = request.Email
	user.UpdatedAt = request.UpdatedAt

	err = uuc.UserRepo.Update(&user)
	return user, err
}

// Xoa project
func (uuc *userUsecase) DeleteUser(id uint) error {
	return uuc.UserRepo.Delete(id)
}

// Lay (Xem chi tiet) project theo ID
func (uuc *userUsecase) GetUserByID(id uint) (entities.User, error) {
	return uuc.UserRepo.GetByID(id)
}

// DI
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepo: repo}
}
