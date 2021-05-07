package user

import "todoAPIGolang/entity"

type UserFormat struct {
	ID        int    `json:"id"`
	FIrstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
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
