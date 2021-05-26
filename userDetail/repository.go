package userDetail

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetByUserID(userID string) (entity.UserDetail, error)
	GetByUserDetailID(UserDetailID string) (entity.UserDetail, error)
	Create(input entity.UserDetail) (entity.UserDetail, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.UserDetail, error)
	// Complete(TodoID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
