package services

import (
	"time"

	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/pkg/encryption"
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

	userExist, err := s.repo.CheckUsernameAndEmail(req.Username, req.Email)

	if err != nil {
		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	if userExist > 0 {
		return entity.AlreadyExistErrorWithAdditionalInfo("Username or Email Address has been used.")
	}

	salt, err := encryption.GenerateSalt()

	if err != nil {
		return entity.GeneralErrorWithAdditionalInfo(err.Error())
	}

	hashedPassword, err := encryption.HashPassword(req.Password, salt)

	if err != nil {
		return entity.GeneralErrorWithAdditionalInfo(err.Error())
	}

	u := entity.User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
		Salt: salt,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(u); err != nil {
		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}