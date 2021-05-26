package category

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Category, error)
	FindByID(categoryID string) (entity.Category, error)
	Create(input entity.Category) (entity.Category, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Category, error)
	// DeleteByID(categoryID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *repository) Create(input entity.Category) (entity.Category, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Category, error) {
	var category entity.Category

	if err := r.db.Model(&category).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return category, err
	}

	if err := r.db.Where("id = ?", ID).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil

}

// func (r *repository) DeleteByID(categoryID string) (string, error) {
// 	if err := r.db.Where("id = ?", categoryID).Delete(&entity.Category{}).Error; err != nil {
// 		return "error", err
// 	}

// 	status := fmt.Sprintf("category id %v success deleted", categoryID)

// 	return status, nil
// }

func (r *repository) FindByID(categoryID string) (entity.Category, error) {
	var category entity.Category

	if err := r.db.Where("user_id = ?", categoryID).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}
