package services

import (
	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/repositories"
)

type InterfaceUserService interface {
	CreateUser(req *entity.CreateUserRequest) *entity.CustomError
}

type UserService struct {
	repo	repositories.InterfaceUserRepository
}

func NewUserService(repo repositories.InterfaceUserRepository) InterfaceUserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *entity.CreateUserRequest) *entity.CustomError {
	return nil
}