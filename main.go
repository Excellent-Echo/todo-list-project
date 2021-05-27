package main

import (
	"todoAPIGolang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// read static file in server app
	r.Static("/images", "./images")

	routes.UserRoute(r)
	routes.UserDetailRoute(r)
	routes.UserProfileRoute(r)
	routes.TodoRoute(r)
	routes.CategoryRoute(r)

	r.Run()
}
