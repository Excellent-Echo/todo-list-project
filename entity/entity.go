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
	DeletedAt   time.Time `gorm:"index"`
}

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	// Todos       []Todo `gorm:"foreignKey:CategoryID"`
}

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	NoHandphone uint
	Gender      string
	Address     string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}
