package handler

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/todo"
)

type todoHandler struct {
	todoService todo.Service
	authService auth.Service
}

func NewTodoHandler(todoService todo.Service, authService auth.Service) *todoHandler {
	return &todoHandler{todoService, authService}
}
