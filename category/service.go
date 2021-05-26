package category

import (
	"errors"
	"fmt"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"
)

type Service interface {
	GetAllCategory() ([]entity.Category, error)
	SaveNewCategory(input entity.CategoryInput) (entity.Category, error)
	UpdateCategoryByID(categoryID string, dataInput entity.UpdateCategoryInput) (entity.Category, error)
	// DeleteCategoryByID(categoryID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllCategory() ([]entity.Category, error) {
	categories, err := s.repository.FindAll()

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *service) SaveNewCategory(input entity.CategoryInput) (entity.Category, error) {

	var newCategory = entity.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	craeateCategory, err := s.repository.Create(newCategory)

	if err != nil {
		return craeateCategory, err
	}

	return craeateCategory, nil
}

func (s *service) UpdateCategoryByID(categoryID string, dataInput entity.UpdateCategoryInput) (entity.Category, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(categoryID); err != nil {
		return entity.Category{}, err
	}

	todo, err := s.repository.FindByID(categoryID)

	if err != nil {
		return entity.Category{}, err
	}

	if todo.ID == 0 {
		newError := fmt.Sprintf("todo id %s not found", categoryID)
		return entity.Category{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}
	// fmt.Println(dataUpdate)

	catUpdated, err := s.repository.UpdateByID(categoryID, dataUpdate)

	if err != nil {
		return catUpdated, err
	}

	return catUpdated, nil
}

// func (s *service) DeleteCategoryByID(categoryID string) (interface{}, error) {
// 	if err := helper.ValidateIDNumber(categoryID); err != nil {
// 		return nil, err
// 	}

// 	cat, err := s.repository.FindByID(categoryID)

// 	if err != nil {
// 		return nil, err
// 	}
// 	if cat.ID == 0 {
// 		newError := fmt.Sprintf("category id %s not found", categoryID)
// 		return nil, errors.New(newError)
// 	}

// 	status, err := s.repository.DeleteByID(categoryID)

// 	if status == "error" {
// 		return nil, errors.New("error delete in internal server")
// 	}

// 	msg := fmt.Sprintf("success delete category ID : %s", categoryID)

// 	formatDelete := FormatDeleteCategory(msg)

// 	return formatDelete, nil
// }
