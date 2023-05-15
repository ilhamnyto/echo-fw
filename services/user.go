package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/pkg/encryption"
	"github.com/ilhamnyto/echo-fw/pkg/token"
	"github.com/ilhamnyto/echo-fw/repositories"
)

type InterfaceUserService interface {
	CreateUser(req *entity.CreateUserRequest) *entity.CustomError
	Login(req *entity.UserLoginRequest) (*entity.UserLoginResponseData, *entity.CustomError)
	GetAllUser(cursor string) ([]*entity.UserData, *entity.Paging, *entity.CustomError)
	GetUserByUsername(username string) (*entity.UserData, *entity.CustomError)
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

func (s *UserService) Login(req *entity.UserLoginRequest) (*entity.UserLoginResponseData, *entity.CustomError) {
	auth, err := s.repo.GetUserCredentialByUsername(req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundErrorWithAdditionalInfo("Wrong username.")
		}
		
		return nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}
	if err = encryption.ValidatePassword(auth.Password, req.Password, auth.Salt); err != nil {
		return nil, entity.GeneralErrorWithAdditionalInfo("Wrong Password.")
	}

	token, err := token.GenerateToken(auth.ID)
	
	if err != nil {
		return nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
	}

	return &entity.UserLoginResponseData{
		Token: token,
	}, nil
}

func (s *UserService) GetAllUser(cursor string) ([]*entity.UserData, *entity.Paging, *entity.CustomError) {
	var (
		t time.Time
		users []*entity.User
		err error
	)

	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		
		if err != nil {
			return nil, nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
		}
		timestamp := int64(cursorInt)
		t = time.Unix(timestamp, 0)
		users, err = s.repo.GetAllUser(&t)
	}else{
		users, err = s.repo.GetAllUser(nil)
	}


	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, entity.NotFoundError()
		}

		return nil, nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var usersData []*entity.UserData
	if len(users) == 6 {
		for _, user := range users[:5] {
			tempUser := new(entity.UserData)
			tempUser.ParseEntityToResponse(user)
			usersData = append(usersData, tempUser)		
		}

	}else {
		for _, user := range users {
			tempUser := new(entity.UserData)
			tempUser.ParseEntityToResponse(user)
			usersData = append(usersData, tempUser)		
		}

	}


	paging := entity.Paging{}

	if len(users) == 6 {
		paging.Next = true
		paging.Cursor = strconv.Itoa(int(usersData[len(usersData)-1].CreatedAt.Unix()))
		fmt.Print(paging)
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return usersData, &paging, nil
}

func (s *UserService) GetUserByUsername(username string) (*entity.UserData, *entity.CustomError) {
	user, err := s.repo.GetUserByUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundError()
		}

		return nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	userData := entity.UserData{}
	userData.ParseEntityToResponse(user)

	return &userData, nil
}