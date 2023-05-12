package routes

import (
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo, c controller.PostController) {
	e.GET("/posts", c.CreatePost)
}