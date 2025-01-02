package entities

import (
	"time"
)

type Project struct {
	ID                uint          `gorm:"primaryKey;autoIncrement"`
	Name              string        `gorm:"size:255;not null"`
	Category          string        `gorm:"not null;default:client"`
	ProjectSpend      int           `gorm:"not null;default:0"`
	ProjectVariance   int           `gorm:"not null;default:0"`
	RevenueRecognised int           `gorm:"not null;default:0"`
	ProjectStartedAt  time.Time     `gorm:"not null"`
	ProjectEndedAt    time.Time
	CreatedAt         time.Time     `gorm:"autoCreateTime"`
	UpdatedAt         time.Time     `gorm:"autoUpdateTime"`
	DeletedAt         time.Time     `gorm:"index"`
	Users             []User        `gorm:"many2many:project_users"`
}
