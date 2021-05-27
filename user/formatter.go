package user

import (
	"time"
	"todoAPIGolang/entity"
)

type UserFormat struct {
	ID        int    `json:"id"`
	FIrstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDetailFormat struct {
	UserFormat
	UserProfile entity.UserProfile `json:"user_profile"`
	UserDetail  entity.UserDetail  `json:"user_detail"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		FIrstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return formatUser
}

func FormatDetailUser(user entity.UserDetailOutput) UserDetailFormat {
	var formatUser = UserDetailFormat{}

	formatUser.ID = user.ID
	formatUser.FIrstName = user.FirstName
	formatUser.LastName = user.LastName
	formatUser.Email = user.Email
	formatUser.UserProfile = user.UserProfile
	formatUser.UserDetail = user.UserDetail

	return formatUser
}

func FormatDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
