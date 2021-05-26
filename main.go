package main

import (
	"todoAPIGolang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.UserDetailRoute(r)
	routes.TodoRoute(r)
	routes.CategoryRoute(r)

	r.Run()
}
