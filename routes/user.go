package routes

import (
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/ilhamnyto/echo-fw/middleware"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, c *controller.UserController) {
	var (
		authGroup = e.Group("/api/v1/auth")
		usersGroup = e.Group("/api/v1/users", middleware.ValidateAuth)
	)

	authGroup.POST("/register", c.CreateUser)
	authGroup.POST("/login", c.Login)
	usersGroup.GET("", c.GetAllUser)
	usersGroup.GET("/:username", c.GetUserByUsername)
	usersGroup.GET("/search", c.SearchUser)
	usersGroup.GET("/me", c.UserProfile)
	usersGroup.PUT("/profile/update", c.UpdateProfile)
	usersGroup.PUT("/profile/update_password", c.UpdatePassword)
}