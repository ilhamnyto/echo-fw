package routes

import (
	"github.com/ilhamnyto/echo-fw/controller"
	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Echo, c controller.PostController) {
	postsGroup := e.Group("/api/v1/posts")
	postsGroup.GET("", c.GetAllPost)
	postsGroup.POST("/create", c.CreatePost)
	postsGroup.GET("/:postid", c.GetPost)
	postsGroup.GET("/search", c.SearchPost)
	postsGroup.GET("/me", c.MyPost)
}