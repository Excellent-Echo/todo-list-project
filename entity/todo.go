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

type TodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

type UpdateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}
