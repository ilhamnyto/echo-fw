package controller

import (
	"github.com/ilhamnyto/echo-fw/entity"
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
	req := entity.CreatePostRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if custErr := p.service.CreatePost(&req, 1); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.CreatedSuccess()
	return c.JSON(resp.StatusCode, resp)
}

func (p *PostController) GetAllPost(c echo.Context) error {
	cursor := c.QueryParam("cursor")
	posts, paging, custErr := p.service.GetAllPost(cursor)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	resp := entity.DataResponse{Data: posts, Paging: paging}

	return c.JSON(200, resp)
}

func (p *PostController) GetPost(c echo.Context) error {
	path := c.Param("postid")
	cursor := c.QueryParam("cursor")
	
	if len(path) > 0 && path[0] == '@' {
		posts, paging, custErr := p.service.GetUserPost(path[1:], cursor)

		if custErr != nil {
			return c.JSON(custErr.StatusCode, custErr)
		}

		resp := entity.DataResponse{Data: posts, Paging: paging}

		return c.JSON(200, resp)
		
	} else {
		post, custErr := p.service.GetPost(path)

		if custErr != nil {
			return c.JSON(custErr.StatusCode, custErr)
		}
	
		resp := entity.DataResponse{Data: post}
	
		return c.JSON(200, resp)
		
	}
}

func (p *PostController) SearchPost(c echo.Context) error {
	query := c.QueryParam("query")
	cursor := c.QueryParam("cursor")

	posts, paging, custErr := p.service.GetUserPostByUsernameOrBody(query, cursor)

		if custErr != nil {
			return c.JSON(custErr.StatusCode, custErr)
		}

		resp := entity.DataResponse{Data: posts, Paging: paging}

		return c.JSON(200, resp)
}

func (p *PostController) MyPost(c echo.Context) error {
	cursor := c.QueryParam("cursor")

	posts, paging, custErr := p.service.GetMyPost(2, cursor)

		if custErr != nil {
			return c.JSON(custErr.StatusCode, custErr)
		}

		resp := entity.DataResponse{Data: posts, Paging: paging}

		return c.JSON(200, resp)
}