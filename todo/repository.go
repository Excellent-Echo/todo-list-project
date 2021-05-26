package todo

import (
	"fmt"
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Todo, error)
	FindAllByUserID(userID string) ([]entity.Todo, error)
	FindByID(TodoID string) (entity.Todo, error)
	Create(todo entity.Todo) (entity.Todo, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Todo, error)
	DeleteByID(TodoID string) (string, error)
	// Complete(TodoID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Todo, error) {
	var todos []entity.Todo

	if err := r.db.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) FindAllByUserID(userID string) ([]entity.Todo, error) {
	var todos []entity.Todo

	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) FindByID(TodoID string) (entity.Todo, error) {
	var todo entity.Todo

	if err := r.db.Where("user_id = ?", TodoID).Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil

}
func (r *repository) Create(todo entity.Todo) (entity.Todo, error) {

	if err := r.db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil

}
func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Todo, error) {
	var todo entity.Todo

	if err := r.db.Model(&todo).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return todo, err
	}

	if err := r.db.Where("id = ?", ID).Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil

}
func (r *repository) DeleteByID(TodoID string) (string, error) {
	if err := r.db.Where("id = ?", TodoID).Delete(&entity.Todo{}).Error; err != nil {
		return "error", err
	}

	status := fmt.Sprintf("todo id %v success deleted", TodoID)

	return status, nil
}
