package main

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/config"
	"todoAPIGolang/handler"
	"todoAPIGolang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowUserHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)

	// r.POST("/users/login", userHandler.CreateUserHandler)

	r.Run(":4444")
}
