package entities

import "time"

type User struct {
	ID        uint          `gorm:"primaryKey;autoIncrement"`
	Name      string        `gorm:"size:255;not null"`
	Email     string        `gorm:"unique;not null"`
	Password  string        `gorm:"not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	DeletedAt time.Time     `gorm:"index"`
	Projects  []Project     `gorm:"many2many:project_users"`
}

/*
 * Cho du truong trong entities co cho phep gia tri null, cung khong duoc phep dat con tro tuy tien
 * Neu 1 truong la 1 mang hoac 1 list hoac 1 map, nen khai bao kieu: DataField []DataType
*/
