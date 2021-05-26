package category

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Category, error)
	Create(input entity.Category) (entity.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
