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

func (u *UserController) Login(c echo.Context) error {
	req := entity.UserLoginRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	token, custErr := u.service.Login(&req)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: token}

	return c.JSON(201, resp)
}

func (u *UserController) GetAllUser(c echo.Context) error {
	cursor := c.QueryParam("cursor")
	users, paging, custErr := u.service.GetAllUser(cursor)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: users, Paging: paging}
	
	return c.JSON(200, resp)
}

func (u *UserController) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")

	userData, custErr := u.service.GetUserByUsername(username)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: userData}
	
	return c.JSON(200, resp)
}