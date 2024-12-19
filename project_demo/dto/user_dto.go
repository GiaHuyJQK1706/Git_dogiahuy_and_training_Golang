package dto

import (
	"time"
)

// CreateUserRequest - Dữ liệu yêu cầu để tạo User
type CreateUserRequest struct {
	Name       string     `json:"name" binding:"required,max=255"`  // Tên bắt buộc, tối đa 255 ký tự
	Email      string     `json:"email" binding:"required,email"`   // Email phải đúng định dạng
	Password   string     `json:"password" binding:"required,min=8,max=20"` // Mật khẩu 8-20 ký tự
	CreatedAt  time.Time  `json:"create_at"`
}

// UpdateUserRequest - Dữ liệu yêu cầu để cập nhật User
type UpdateUserRequest struct {
	Name       string     `json:"name" binding:"omitempty,max=255"` // Tên không bắt buộc
	Email      string     `json:"email" binding:"omitempty,email"`  // Email không bắt buộc
	UpdatedAt  time.Time  `json:"create_at"`
}
