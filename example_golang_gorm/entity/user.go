package entity //Bai 1

import "time"

// User entity for users table.
type User struct {
	ID        int64     `gorm:"primaryKey"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Projects  []Project `gorm:"foreignKey:UserID"`
}
