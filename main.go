package main

import (
	"todoAPIGolang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.GetAllUser)
	r.GET("/users/:user_id", handler.GetUserByID)
	r.POST("/users", handler.CreateNewUser)
	r.PUT("/users/:user_id", handler.UpdateUserByID)
	r.DELETE("/users/:user_id", handler.DeleteByUserID)
	r.Run(":4444")
}
