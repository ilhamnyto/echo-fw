package services

import (
	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/repositories"
)

type InterfacePostService interface {
	CreatePost(req *entity.CreatePostRequest) *entity.CustomError
}

type PostService struct {
	repo repositories.InterfacePostRepository
}

func NewPostService(repo repositories.InterfacePostRepository) InterfacePostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(req *entity.CreatePostRequest) *entity.CustomError {
	return nil
}