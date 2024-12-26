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

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) Create(project *entities.Project) error {
	args := m.Called(project)
	
	return args.Error(0)
}

func (m *MockProjectRepository) Update(project *entities.Project) error {
	args := m.Called(project)

	return args.Error(0)
}

func (m *MockProjectRepository) Delete(id uint) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockProjectRepository) GetByID(id uint) (entities.Project, error) {
	args := m.Called(id)

	return args.Get(0).(entities.Project), args.Error(1)
}

func (m *MockProjectRepository) List() ([]entities.Project, error) {
	args := m.Called()

	return args.Get(0).([]entities.Project), args.Error(1)
}

func (m *MockProjectRepository) GetProjectsByUserID(userID uint) ([]entities.Project, error) {
	args := m.Called(userID)

	return args.Get(0).([]entities.Project), args.Error(1)
}

func TestCreateProjectUsecase(t *testing.T) {
	mockRepo := new(MockProjectRepository)
	uc := usecase.NewProjectUsecase(mockRepo)

	request := dto.CreateProjectRequest{Name: "Test Project"}
	mockRepo.On("Create", mock.Anything).Return(nil)

	project, err := uc.CreateProject(request)

	assert.Nil(t, err)
	assert.Equal(t, "Test Project", project.Name)
	mockRepo.AssertCalled(t, "Create", mock.Anything)
}
