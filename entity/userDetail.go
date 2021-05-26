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
