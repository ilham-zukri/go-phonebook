package services

import (
	"github.com/ilham-zukri/go-phonebook/models"
	"github.com/ilham-zukri/go-phonebook/repositories"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Register(user *models.User) error {
	panic("unimplemented")
}
