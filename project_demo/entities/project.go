package entities

import (
	"time"
)

type Project struct {
	ID                uint      `gorm:"primaryKey"`
	Name              string    `gorm:"size:255;not null"`
	Category          string    `gorm:"not null;default:'client'"`
	ProjectSpend      int       `gorm:"not null;default:0"`
	ProjectVariance   int       `gorm:"not null;default:0"`
	RevenueRecognised int       `gorm:"not null;default:0"`
	ProjectStartedAt  time.Time `gorm:"not null"`
	ProjectEndedAt    *time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
