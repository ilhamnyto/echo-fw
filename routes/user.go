package routes

import (
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, c *controller.UserController) {
	var (
		authGroup = e.Group("/api/v1/auth")
		usersGroup = e.Group("/api/v1/users")
	)

	authGroup.POST("/register", c.CreateUser)
	usersGroup.GET("/", c.CreateUser)
}