package todo

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Todo, error)
	GetAllByUserID(userID string) ([]entity.Todo, error)
	GetByTodoID(TodoID string) (entity.Todo, error)
	Create(input entity.Todo) (entity.Todo, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Todo, error)
	Delete(TodoID string) (string, error)
	// Complete(TodoID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.Todo, error) {
	var todos []entity.Todo

	return todos, nil
}
