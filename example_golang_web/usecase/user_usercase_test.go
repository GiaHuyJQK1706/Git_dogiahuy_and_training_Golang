package usecase

import (
	"example_golang_web/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Insert(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{
		FirstName: "Do",
		LastName:  "Huy",
		Email:     "dohuy@example.com",
		Password:  "password123",
	}

	mockRepo.On("Insert", mock.Anything).Return(user, nil)

	createdUser, err := usecase.CreateUser(user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "Do", createdUser.FirstName)
	assert.Equal(t, "Huy", createdUser.LastName)
	mockRepo.AssertExpectations(t)
}