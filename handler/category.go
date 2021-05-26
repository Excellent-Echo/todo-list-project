package handler

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/category"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"

	"github.com/gin-gonic/gin"
)

type categoryhandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(categoryService category.Service, authService auth.Service) *categoryhandler {
	return &categoryhandler{categoryService, authService}
}

func (h *categoryhandler) GetAllCategoryhandler(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategory()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all category", 200, "success", categories)
	c.JSON(200, response)

}

func (h *categoryhandler) SaveNewCategoryHandler(c *gin.Context) {

	var inputCategory entity.CategoryInput

	if err := c.ShouldBindJSON(&inputCategory); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newCategory, err := h.categoryService.SaveNewCategory(inputCategory)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new category", 201, "success", newCategory)
	c.JSON(201, response)
}

func (h *categoryhandler) UpdateCategoryByIDHandler(c *gin.Context) {

	id := c.Params.ByName("category_id")

	var updateInputCategory entity.UpdateCategoryInput

	if err := c.ShouldBindJSON(&updateInputCategory); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	updateCategory, err := h.categoryService.UpdateCategoryByID(id, updateInputCategory)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update category by ID", 200, "success", updateCategory)
	c.JSON(200, response)
}
