package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" binding:"required"`
	Email     string         `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	Password  string         `json:"password" gorm:"not null"`
	Age       int            `json:"age" binding:"gte=0,lte=130" `
	CreatedAt time.Time      `json:"created_at"`              // Auto-managed
	UpdatedAt time.Time      `json:"updated_at"`              // Auto-managed
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // Auto-managed
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,gte=0,lte=130" `
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" binding:"omitempty,min=2"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
	Age      int    `json:"age,omitempty" binding:"omitempty,gte=0,lte=130"`
}
