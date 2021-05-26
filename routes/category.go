package routes

import (
	"todoAPIGolang/category"
	"todoAPIGolang/handler"

	"github.com/gin-gonic/gin"
)

var (
	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository)
	categoryhandler    = handler.NewCategoryHandler(categoryService, authService)
)

func CategoryRoute(r *gin.Engine) {
	r.GET("/categories", categoryhandler.GetAllCategoryhandler)
	r.POST("/categories", categoryhandler.SaveNewCategoryHandler)
	r.PUT("/categories/:category_id", categoryhandler.UpdateCategoryByIDHandler)
}
