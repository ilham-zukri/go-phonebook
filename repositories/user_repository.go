package repositories

import (
	"github.com/ilham-zukri/go-phonebook/config"
	"github.com/ilham-zukri/go-phonebook/models"
)

type UserRepository interface{
	Create(user *models.User) error
}

type userRepository struct {
	
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error{
	return config.DB.Create(&user).Error
}