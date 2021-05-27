package userProfile

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUserID(userID string) (entity.UserProfile, error)
	Create(userProfile entity.UserProfile) (entity.UserProfile, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.UserProfile, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserID(userID string) (entity.UserProfile, error) {
	var userProfile entity.UserProfile

	if err := r.db.Where("user_id = ?", userID).Find(&userProfile).Error; err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (r *repository) Create(userProfile entity.UserProfile) (entity.UserProfile, error) {
	if err := r.db.Create(&userProfile).Error; err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.UserProfile, error) {
	var userProfile entity.UserProfile

	if err := r.db.Model(&userProfile).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return userProfile, err
	}

	if err := r.db.Where("id = ?", ID).Find(&userProfile).Error; err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
