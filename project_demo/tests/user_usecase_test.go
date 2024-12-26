package tests

// test cho usecase/project_usecase.go

import (
	"project_demo/dto"
	"project_demo/entities"
	"project_demo/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *entities.User) error {
	args := m.Called(user)
	
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *entities.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (entities.User, error) {
	args := m.Called(id)

	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (entities.User, error) {
	args := m.Called(email)

	return args.Get(0).(entities.User), args.Error(1)
}

func TestCreateUserUsecase(t *testing.T) {
	mockRepo := new(MockUserRepository)
	uc := usecase.NewUserUsecase(mockRepo)

	request := dto.CreateUserRequest{Name: "Test User"}
	mockRepo.On("Create", mock.Anything).Return(nil)

	user, err := uc.CreateUser(request)

	assert.Nil(t, err)
	assert.Equal(t, "Test User", user.Name)
	mockRepo.AssertCalled(t, "Create", mock.Anything)
}
