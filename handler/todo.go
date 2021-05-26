package handler

import (
	"strconv"
	"todoAPIGolang/auth"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"
	"todoAPIGolang/todo"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService todo.Service
	authService auth.Service
}

func NewTodoHandler(todoService todo.Service, authService auth.Service) *todoHandler {
	return &todoHandler{todoService, authService}
}

func (h *todoHandler) GetAllTodoHandler(c *gin.Context) {
	todos, err := h.todoService.GetAllTodo()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all todo", 200, "success", todos)
	c.JSON(200, response)
}

func (h *todoHandler) GetAllTodoByUserIDHandler(c *gin.Context) {

	// id := c.Params.ByName("user_id")
	userData := int(c.MustGet("currentUser").(int))

	userID := strconv.Itoa(userData)

	todos, err := h.todoService.GetAllTodoByUserID(userID)

	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get all todo by user ID", 200, "success", todos)
	c.JSON(200, response)

}

func (h *todoHandler) GetTodoByIDHandler(c *gin.Context) {
	id := c.Params.ByName("todo_id")

	todo, err := h.todoService.GetTodoByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get todo by id", 200, "status OK", todo)
	c.JSON(200, response)
}

func (h *todoHandler) CreateNewTodoHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	var inputTodo entity.TodoInput

	if err := c.ShouldBindJSON(&inputTodo); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newTodo, err := h.todoService.SaveNewTodo(inputTodo, userID)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new todo", 201, "success", newTodo)
	c.JSON(201, response)
}

func (h *todoHandler) UpdateTodoByIDHandler(c *gin.Context) {
	id := c.Params.ByName("todo_id")

	var updateTodoInput entity.UpdateTodoInput

	if err := c.ShouldBindJSON(&updateTodoInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	todo, err := h.todoService.GetTodoByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData := int(c.MustGet("currentUser").(int))

	// validate todo user id with userData in localstorage
	if todo.UserID != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.todoService.UpdateTodoByID(id, updateTodoInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update todo by ID", 200, "success", user)
	c.JSON(200, response)
}

func (h *todoHandler) DeleteTodoByIDHandler(c *gin.Context) {
	id := c.Params.ByName("todo_id")

	todo, err := h.todoService.GetTodoByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData := int(c.MustGet("currentUser").(int))

	// validate todo user id with userData in localstorage
	if todo.UserID != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	deleteTodo, err := h.todoService.DeleteTodoByID(id)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success delete todo by ID", 200, "success", deleteTodo)
	c.JSON(200, response)

}

func (h *todoHandler) CompleteTodoHandler(c *gin.Context) {
	id := c.Params.ByName("todo_id")

	todo, err := h.todoService.GetTodoByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData := int(c.MustGet("currentUser").(int))

	// validate todo user id with userData in localstorage
	if todo.UserID != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	competeTodo, err := h.todoService.CompleteTodo(id)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success complete todo by ID", 200, "success", competeTodo)
	c.JSON(200, response)
}
