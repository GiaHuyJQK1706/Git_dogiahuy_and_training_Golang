package dto

import (
	"time"
)

type CreateProjectRequest struct {
	Name             string     `json:"name" binding:"required,max=255"`
	Category         string     `json:"category"`
	ProjectSpend     int        `json:"project_spend"`
	ProjectVariance  int        `json:"project_variance"`
	ProjectStartedAt time.Time  `json:"project_started_at" binding:"required"`
	ProjectEndedAt   *time.Time  `json:"project_ended_at"`
}

type CreateProjectResponse struct {
	Name             string     `json:"name"`
	Category         string     `json:"category"`
	ProjectSpend     int        `json:"project_spend"`
	ProjectVariance  int        `json:"project_variance"`
	ProjectStartedAt time.Time  `json:"project_started_at"`
	ProjectEndedAt   *time.Time `json:"project_ended_at"`
}

/*
 * Khi truong nao co the la null thi se dat con tro
 * Khi truong nao khong the la null thi khong duoc dat con tro
*/
