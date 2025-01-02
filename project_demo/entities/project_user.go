package entities

import "time"

type ProjectUser struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ProjectID uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"index"`
}
