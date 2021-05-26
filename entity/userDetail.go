package entity

import "time"

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	NoHandphone uint
	Gender      string
	Address     string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   time.Time `gorm:"index"`
}

type UserDetailInput struct {
	NoHandphone uint   `json:"no_handphone" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type UpdateUserDetailInput struct {
	NoHandphone uint   `json:"no_handphone"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
}
