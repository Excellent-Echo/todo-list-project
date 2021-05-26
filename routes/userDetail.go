package routes

import (
	"todoAPIGolang/handler"
	"todoAPIGolang/userDetail"

	"github.com/gin-gonic/gin"
)

var (
	userDetailRepository = userDetail.NewRepository(DB)
	userDetailService    = userDetail.NewService(userDetailRepository)
	userDetailHanlder    = handler.NewUserDetailHandler(userDetailService, authService)
)

func UserDetailRoute(r *gin.Engine) {
	r.GET("/user_details", handler.Middleware(userService, authService), userDetailHanlder.GetUserDetailByUserIDHandler)
	r.POST("/user_details", handler.Middleware(userService, authService), userDetailHanlder.SaveNewUserDetailHandler)
	r.PUT("/user_details", handler.Middleware(userService, authService), userDetailHanlder.UpdateUserDetailByIDHandler)
}
