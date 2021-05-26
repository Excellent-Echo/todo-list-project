package entity

import "time"

type UserDetail struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	NoHandphone uint      `json:"no_handphone"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
