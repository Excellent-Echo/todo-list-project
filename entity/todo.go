package entity

import "time"

type Todo struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	CategoryID  int
	UserID      int
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt   time.Time `gorm:"index"`
}
