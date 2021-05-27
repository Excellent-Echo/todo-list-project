package userProfile

import (
	"errors"
	"fmt"
	"strconv"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"
)

type Service interface {
	GetUserProfileByUserID(userID string) (entity.UserProfile, error)
	SavenewUserProfile(pathFile string, userID int) (entity.UserProfile, error)
	UpdateUserProfileByID(pathFile string, userProfileID string) (entity.UserProfile, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserProfileByUserID(userID string) (entity.UserProfile, error) {
	userProfile, err := s.repository.FindByUserID(userID)

	if err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (s *service) SavenewUserProfile(pathFile string, userID int) (entity.UserProfile, error) {

	ID := strconv.Itoa(userID)

	userProfileCheck, _ := s.repository.FindByUserID(ID)

	if userProfileCheck.UserID == userID {
		errooStatus := fmt.Sprintf("user profile for user id %d has been created", userID)
		return userProfileCheck, errors.New(errooStatus)
	}

	newUserProfile := entity.UserProfile{
		ProfileUser: pathFile,
		UserID:      userID,
	}

	userProfile, err := s.repository.Create(newUserProfile)

	if err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (s *service) UpdateUserProfileByID(pathFile string, userProfileID string) (entity.UserProfile, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userProfileID); err != nil {
		return entity.UserProfile{}, err
	}

	dataUpdate["profile_user"] = pathFile

	// fmt.Println(dataUpdate)

	userProfileUpdate, err := s.repository.UpdateByID(userProfileID, dataUpdate)

	if err != nil {
		return userProfileUpdate, err
	}

	return userProfileUpdate, nil
}
