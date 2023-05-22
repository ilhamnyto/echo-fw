package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service 	services.InterfaceUserService
	cache *redis.Client
}

func NewUserController(service services.InterfaceUserService, redis *redis.Client) *UserController {
	return &UserController{service: service, cache: redis}
}

func (u *UserController) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
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

func (u *UserController) SearchUser(c echo.Context) error {
	query := c.QueryParam("query")
	cursor := c.QueryParam("cursor")

	users, paging, custErr := u.service.SearchUserByUsernameOrEmail(query, cursor)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: users, Paging: paging}

	return c.JSON(200, resp)
}

func (u *UserController) UserProfile(c echo.Context) error {
	userId := c.Get("user_id").(int)

	result, err := u.cache.Get(c.Request().Context(), fmt.Sprint(userId)).Result()

	if err == nil {
		
		var dr entity.DataResponse

		err := json.Unmarshal([]byte(result), &dr)

		if err != nil {
			return c.JSON(500, err.Error())
		}

		return c.JSON(200, dr)

	}

	userData, custErr := u.service.GetProfile(userId)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: userData}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(500, err.Error())
	}


	err = u.cache.Set(c.Request().Context(), fmt.Sprint(userId), string(jsonData), 2 * 60 * time.Second).Err()

	if err != nil {
		return c.JSON(500, "")
	}

	return c.JSON(200, resp)
}

func (u *UserController) UpdateProfile(c echo.Context) error {
	req := entity.UpdateUserRequest{}
	userId := c.Get("user_id").(int)

	if err := c.Bind(&req); err != nil {
		return err
	}

	if custErr := u.service.UpdateUserProfile(&req, userId); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.GeneralSuccessWithCustomMessageAndPayload("Profile updated successfully.", nil)
	
	return c.JSON(resp.StatusCode, resp)
}

func (u *UserController) UpdatePassword(c echo.Context) error {
	req := entity.UpdatePasswordRequest{}
	userId := c.Get("user_id").(int)

	if err := c.Bind(&req); err != nil {
		return err
	}

	if custErr := u.service.UpdateUserPassword(&req, userId); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.GeneralSuccessWithCustomMessageAndPayload("Password updated succesfully.", nil)
	return c.JSON(resp.StatusCode, resp)
}