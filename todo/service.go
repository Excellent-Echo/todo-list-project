package todo

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"
)

type Service interface {
	GetAllTodo() ([]entity.Todo, error)
	GetAllTodoByUserID(userID string) ([]entity.Todo, error)
	GetTodoByID(todoID string) (entity.Todo, error)
	SaveNewTodo(input entity.TodoInput, CategoryID string, UserID string) (entity.Todo, error)
	UpdateTodoByID(todoID string, dataInput entity.TodoInput) (entity.Todo, error)
	DeleteTodoByID(todoID string) (interface{}, error)
	CompleteTodo(todoID string) (entity.Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllTodo() ([]entity.Todo, error) {
	todos, err := s.repository.FindAll()

	if err != nil {
		return todos, err
	}

	return todos, nil
}
func (s *service) GetAllTodoByUserID(userID string) ([]entity.Todo, error) {

	todos, err := s.repository.FindAllByUserID(userID)

	if err != nil {
		return todos, err
	}

	return todos, nil

}
func (s *service) GetTodoByID(todoID string) (entity.Todo, error) {

	todo, err := s.repository.FindByID(todoID)

	if err != nil {
		return todo, err
	}

	return todo, nil

}
func (s *service) SaveNewTodo(input entity.TodoInput, CategoryID string, UserID string) (entity.Todo, error) {

	IDUser, _ := strconv.Atoi(UserID)
	IDCategory, _ := strconv.Atoi(CategoryID)

	var newTodo = entity.Todo{
		Title:       input.Title,
		Description: input.Description,
		CategoryID:  IDCategory,
		UserID:      IDUser,
		IsComplete:  false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createTodo, err := s.repository.Create(newTodo)

	if err != nil {
		return createTodo, err
	}

	return createTodo, nil

}
func (s *service) UpdateTodoByID(todoID string, dataInput entity.TodoInput) (entity.Todo, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(todoID); err != nil {
		return entity.Todo{}, err
	}

	todo, err := s.repository.FindByID(todoID)

	if err != nil {
		return entity.Todo{}, err
	}

	if todo.ID == 0 {
		newError := fmt.Sprintf("todo id %s not found", todoID)
		return entity.Todo{}, errors.New(newError)
	}

	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	todoUpdated, err := s.repository.UpdateByID(todoID, dataUpdate)

	if err != nil {
		return todoUpdated, err
	}

	return todoUpdated, nil
}

func (s *service) CompleteTodo(todoID string) (entity.Todo, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(todoID); err != nil {
		return entity.Todo{}, err
	}

	todo, err := s.repository.FindByID(todoID)

	if err != nil {
		return entity.Todo{}, err
	}

	if todo.ID == 0 {
		newError := fmt.Sprintf("todo id %s not found", todoID)
		return entity.Todo{}, errors.New(newError)
	}

	dataUpdate["is_complete"] = true
	dataUpdate["updated_at"] = time.Now()

	todoUpdated, err := s.repository.UpdateByID(todoID, dataUpdate)

	if err != nil {
		return todoUpdated, err
	}

	return todoUpdated, nil

}

func (s *service) DeleteTodoByID(todoID string) (interface{}, error) {

	if err := helper.ValidateIDNumber(todoID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(todoID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("todo id %s not found", todoID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(todoID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete todo ID : %s", todoID)

	formatDelete := FormatDeleteTodo(msg)

	return formatDelete, nil
}
