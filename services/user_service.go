package services

import (
	"errors"

	"github.com/ilham-zukri/go-phonebook/models"
	"github.com/ilham-zukri/go-phonebook/repositories"
	"github.com/ilham-zukri/go-phonebook/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateUser(user *models.User) error {
	exists, _ := u.userRepository.FindByEmpID(user.EmpID)
	if exists.ID != 0 {
		return errors.New("employe already exist")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return u.userRepository.Create(user)
}
