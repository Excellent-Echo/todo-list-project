package handler

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/userDetail"
)

type userDetailHandler struct {
	userDetailService userDetail.Service
	authService       auth.Service
}

func NewUserDetailHandler(userDetailService userDetail.Service, authService auth.Service) *userDetailHandler {
	return &userDetailHandler{userDetailService, authService}
}
