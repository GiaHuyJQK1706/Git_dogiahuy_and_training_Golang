package usecase

import (
	"example_golang_gorm/entity"
	"example_golang_gorm/repository"
	"time"
	"fmt"

	"crypto/md5"
	"encoding/hex"
)

type UserUsecaseInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id int64) (*entity.User, error)
}

type UserUsecase struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserUsecase(repo repository.UserRepositoryInterface) UserUsecaseInterface {
	// DI: Truyen `UserRepositoryInterface` (Phu thuoc) tu ben ngoai UserUsecase.
	return &UserUsecase{UserRepo: repo}
}

func GetCurrentTimeUTC7() (time.Time, error) {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load location: %v", err)
	}
	
	return time.Now().In(location), nil
}

func (u *UserUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	//Bam mat khau
	hasher := md5.New()
	hasher.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hasher.Sum(nil))

	//Cai dat CreateAt va UpdateAt trung voi gio UTC+7
	currentTime, err := GetCurrentTimeUTC7()
	if err != nil {
		return nil, err
	}
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	return u.UserRepo.Insert(user)
}

// Ham lay thong tin user khi co id
func (u *UserUsecase) GetUserByID(id int64) (*entity.User, error) {
	return u.UserRepo.FindByID(id)
}
