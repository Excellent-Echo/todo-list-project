package user

import (
	"errors"
	"fmt"
	"time"
	"todoAPIGolang/entity"
	"todoAPIGolang/helper"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(userID string) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	// melakukan business logic disini
	users, err := s.repository.FindAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) GetUserByID(userID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)

	return formatUser, nil
}

func (s *service) DeleteUserByID(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}

//(UserFormat, error)

func (s *service) UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}
	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	dataUpdate["updated_at"] = time.Now()

	fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdated)

	return formatUser, nil
}
