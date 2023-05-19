package middleware

import (
	"strings"

	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/pkg/token"
	"github.com/labstack/echo/v4"
)

func ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(403, entity.UnauthorizedError())
		}
		
		tokenString := strings.Split(auth, "Bearer ")
		if len(tokenString) != 2  {
			return c.JSON(403, entity.UnauthorizedError())
		}
		
		token, err := token.ValidateToken(tokenString[1])
		
		if err != nil {
			return c.JSON(403, entity.UnauthorizedError())
		}

		c.Set("user_id", token.UserID)
		
		return next(c)
	}
}