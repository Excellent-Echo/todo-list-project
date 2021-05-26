package entity

import "time"

type Todo struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	UserID      int       `json:"user_id"`
	IsComplete  bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
