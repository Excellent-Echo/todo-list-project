package handler

import (
	"fmt"
	"strconv"
	"todoAPIGolang/helper"
	"todoAPIGolang/userProfile"

	"github.com/gin-gonic/gin"
)

// 3 endpoint
// 1 GET
// 2 POST / create
// 3 update foto PUT

type userProfileHandler struct {
	service userProfile.Service
}

func NewUserProfileHandler(service userProfile.Service) *userProfileHandler {
	return &userProfileHandler{service}
}

func (h *userProfileHandler) GetUserProfileByUserIDHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))

	userID := strconv.Itoa(userData)

	userProfile, err := h.service.GetUserProfileByUserID(userID)

	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get user profile by user ID", 200, "success", userProfile)
	c.JSON(200, response)
}

func (h *userProfileHandler) SaveNewUserProfileHandler(c *gin.Context) {

	userData := int(c.MustGet("currentUser").(int))

	file, err := c.FormFile("profile") // postman

	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	path := fmt.Sprintf("/images/profile-%d-%s", userData, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userProfile, err := h.service.SavenewUserProfile(path, userData)

	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create user profile", 201, "success", userProfile)
	c.JSON(201, response)
}
