package usecase

import (
	"project_demo/dto"
	"project_demo/entities"
	"project_demo/repository"
)

type ProjectUsecase interface {
	CreateProject(dto.CreateProjectRequest) (entities.Project, error)
}

type projectUsecase struct {
	PrọectRepo repository.ProjectRepository
}

func (puc *projectUsecase) CreateProject(request dto.CreateProjectRequest) (entities.Project, error) {
	project := entities.Project{
		Name:             request.Name,
		Category:         request.Category,
		ProjectSpend:     request.ProjectSpend,
		ProjectVariance:  request.ProjectVariance,
		ProjectStartedAt: *request.ProjectStartedAt,
		ProjectEndedAt:   request.ProjectEndedAt,
	}

	err := puc.PrọectRepo.Create(&project)
	return project, err
}

func NewProjectUsecase(repo repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{PrọectRepo: repo}
}
