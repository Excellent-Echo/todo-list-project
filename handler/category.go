package handler

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/category"
)

type categoryhandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(categoryService category.Service, authService auth.Service) *categoryhandler {
	return &categoryhandler{categoryService, authService}
}
