package routes

import (
	"todoAPIGolang/handler"
	"todoAPIGolang/todo"

	"github.com/gin-gonic/gin"
)

var (
	todoRepository = todo.NewRepository(DB)
	todoService    = todo.NewService(todoRepository, userRepository)
	todoHandler    = handler.NewTodoHandler(todoService, authService)
)

func TodoRoute(r *gin.Engine) {
	r.GET("/todos", handler.Middleware(userService, authService), todoHandler.GetAllTodoHandler)
	r.POST("/todos", handler.Middleware(userService, authService), todoHandler.CreateNewTodoHandler)
	r.GET("/todos/users/:user_id", handler.Middleware(userService, authService), todoHandler.GetAllTodoByUserIDHandler)
	r.GET("/todos/:todo_id", handler.Middleware(userService, authService), todoHandler.GetTodoByIDHandler)
	r.PUT("/todos/:todo_id", handler.Middleware(userService, authService), todoHandler.UpdateTodoByIDHandler)
	r.PATCH("/todos/:todo_id/complete", handler.Middleware(userService, authService), todoHandler.CompleteTodoHandler)
	r.DELETE("/todos/:todo_id", handler.Middleware(userService, authService), todoHandler.DeleteTodoByIDHandler)
}
