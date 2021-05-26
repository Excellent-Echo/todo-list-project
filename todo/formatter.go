package todo

import (
	"time"
)

// type TodoFormat struct {
// 	ID          int       `json:"id"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	IsComplete  bool      `json:"is_complete"`
// 	UserID      int       `json:"user_id"`
// 	CategoryID  int       `json:"category_id"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

// func FormatUser(todo entity.Todo) TodoFormat {
// 	var formatTodo = TodoFormat{
// 		ID:          todo.ID,
// 		Title:       todo.Title,
// 		Description: todo.Description,
// 		UserID:      todo.UserID,
// 		CategoryID:  todo.CategoryID,
// 		CreatedAt:   todo.CreatedAt,
// 		UpdatedAt:   todo.UpdatedAt,
// 	}

// 	return formatTodo
// }

func FormatDeleteTodo(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
