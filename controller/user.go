package controller

import (
	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/services"
	"github.com/ilhamnyto/echo-fw/utils"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service 	services.InterfaceUserService
}

func NewUserController(service services.InterfaceUserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if custErr := utils.ValidateRegisterRequest(&req); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	if custErr := u.service.CreateUser(&req); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.CreatedSuccess()
	return c.JSON(resp.StatusCode, resp)
}