package entity

type UserInput struct {
	FirstName string `json:"first_name" binding:"required,default(stranger)"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}
