package repositories

import (
	"github.com/ilham-zukri/go-phonebook/config"
	"github.com/ilham-zukri/go-phonebook/models"
)

type UserRepository interface{
	Create(user *models.User) error
	FindByEmpID(empID string) (*models.User, error)
}

type userRepository struct {
	
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error{
	return config.DB.Create(&user).Error
}

func (r *userRepository) FindByEmpID (empID string) (*models.User, error){
	var user models.User
	error := config.DB.Where("emp_id = ?", empID).First(&user).Error
	return &user, error
}