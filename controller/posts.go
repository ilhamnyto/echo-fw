package controller

import (

	"github.com/ilhamnyto/echo-fw/services"
	"github.com/labstack/echo/v4"
)

type PostController struct {
	service services.InterfacePostService
}

func NewPostController(service services.InterfacePostService) *PostController {
	return &PostController{service: service}
}

func (p *PostController) CreatePost(c echo.Context) error {
	return c.JSON(200, "hello from post")
}