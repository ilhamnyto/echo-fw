package controller

import (
	"github.com/ilhamnyto/echo-fw/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service 	services.InterfaceUserService
}

func NewUserController(service services.InterfaceUserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) Hello(c echo.Context) error {
	return c.JSON(200, "hello world")
}