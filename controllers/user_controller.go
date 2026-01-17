package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilham-zukri/go-phonebook/models"
	"github.com/ilham-zukri/go-phonebook/services"
	"github.com/ilham-zukri/go-phonebook/utils"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "failed to parse request body", err.Error())
	}

	if err := u.service.CreateUser(user); err != nil {
		return utils.InternalServerError(ctx, "failed to create user", err.Error())
	}

	var userResponse models.UserResponse
	err := copier.Copy(&userResponse, &user)

	if err != nil {
		return utils.InternalServerError(ctx, "failed to copy user", err.Error())
	}

	return utils.Success(ctx, "user created successfully", userResponse)
}
