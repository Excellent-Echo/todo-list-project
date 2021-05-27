package routes

import (
	"todoAPIGolang/handler"
	"todoAPIGolang/userProfile"

	"github.com/gin-gonic/gin"
)

var (
	userProfileRepository = userProfile.NewRepository(DB)
	userProfileService    = userProfile.NewService(userProfileRepository)
	userProfileHandler    = handler.NewUserProfileHandler(userProfileService)
)

func UserProfileRoute(r *gin.Engine) {
	r.GET("/user_profile", handler.Middleware(userService, authService), userProfileHandler.GetUserProfileByUserIDHandler)
	r.POST("/user_profile", handler.Middleware(userService, authService), userProfileHandler.SaveNewUserProfileHandler)
	r.PUT("/user_profile", handler.Middleware(userService, authService), userProfileHandler.UpdateUserProfileByIDHandler) // otomatis terhubung ke user yang login
}
