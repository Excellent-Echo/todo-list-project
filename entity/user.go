package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
	// Todos      []Todo       `gorm:"foreignKey:UserID"`
	// UserDetail []UserDetail `gorm:"foreignKey:UserID"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDetailOutput struct {
	ID          int         `gorm:"primaryKey" json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	Password    string      `json:"-"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	UserProfile UserProfile `json:"user_profile"`
	UserDetail  UserDetail  `json:"user_detail"`
}
