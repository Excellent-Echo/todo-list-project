package user

import (
	"todoAPIGolang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindByID(ID string) (entity.User, error)
	DeleteByID(ID string) (string, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByIDwithUserDetailUserProfile(ID string) (entity.UserDetailOutput, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByIDwithUserDetailUserProfile(ID string) (entity.UserDetailOutput, error) {
	var user entity.UserDetailOutput

	if err := r.db.Where("id = ?", ID).Find(&user).Preload("UserDetail").Preload("UserProfile").Error; err != nil {
		return user, err
	}

	return user, nil
}
