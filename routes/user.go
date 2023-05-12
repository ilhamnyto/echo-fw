package routes

import (
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, c *controller.UserController) {
	e.GET("/", c.Hello)
}