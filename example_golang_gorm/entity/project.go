package entity // Bai 1

import "time"

// Project entity for projects table.
type Project struct {
	ID               int64     `gorm:"primaryKey"`
	Name             string    `gorm:"not null"`
	ProjectStartedAt time.Time `gorm:"not null"`
	ProjectEndedAt   *time.Time
	UserID           int64     `gorm:"not null"`
	User             User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
