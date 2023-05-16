package services

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/pkg/encryption"
	"github.com/ilhamnyto/echo-fw/pkg/token"
	"github.com/ilhamnyto/echo-fw/repositories"
	"github.com/ilhamnyto/echo-fw/utils"
)

type InterfaceUserService interface {
	CreateUser(req *entity.CreateUserRequest) *entity.CustomError
	Login(req *entity.UserLoginRequest) (*entity.UserLoginResponseData, *entity.CustomError)
	GetAllUser(cursor string) ([]*entity.UserData, *entity.Paging, *entity.CustomError)
	GetUserByUsername(username string) (*entity.UserData, *entity.CustomError)
	SearchUserByUsernameOrEmail(query string, cursor string) ([]*entity.UserData, *entity.Paging, *entity.CustomError)
	GetProfile(userID int) (*entity.UserData, *entity.CustomError)
	UpdateUserProfile(req *entity.UpdateUserRequest, userID int) *entity.CustomError
	UpdateUserPassword(req *entity.UpdatePasswordRequest, userID int) *entity.CustomError
}

type UserService struct {
	repo	repositories.InterfaceUserRepository
}

func NewUserService(repo repositories.InterfaceUserRepository) InterfaceUserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *entity.CreateUserRequest) *entity.CustomError {

	if err := utils.ValidateRegisterRequest(req); err != nil {
		return entity.BadRequestErrorWithAdditionalInfo(err.Error())
	}

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

func (s *UserService) SearchUserByUsernameOrEmail(query string, cursor string) ([]*entity.UserData, *entity.Paging, *entity.CustomError) {
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
		users, err = s.repo.GetUserBySimalarUsernameOrEmail(query, &t)
	}else{
		users, err = s.repo.GetUserBySimalarUsernameOrEmail(query, nil)
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
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return usersData, &paging, nil
}

func (s *UserService) GetProfile(userID int) (*entity.UserData, *entity.CustomError) {
	user, err := s.repo.GetUserByUserID(userID)

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

func (s *UserService) UpdateUserProfile(req *entity.UpdateUserRequest, userID int) *entity.CustomError {
	if err := utils.ValidateUpdateUserProfileRequest(req); err != nil {
		return entity.BadRequestErrorWithAdditionalInfo(err.Error())
	}

	if err := s.repo.UpdateUser(req, 1); err != nil {
		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (s *UserService) UpdateUserPassword(req *entity.UpdatePasswordRequest, userID int) *entity.CustomError {
	if err := utils.ValidateUpdatePasswordRequest(req); err != nil {
		return entity.BadRequestErrorWithAdditionalInfo(err.Error())
	}

	user, err := s.repo.GetUserPasswordAndSalt(userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.NotFoundError()
		}

		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	if err := encryption.ValidatePassword(user.Password, req.Password, user.Salt); err == nil {
		return entity.BadRequestErrorWithAdditionalInfo("Password cant be the same as before.")
	}

	salt, err := encryption.GenerateSalt()

	if err != nil {
		return entity.GeneralErrorWithAdditionalInfo(err.Error())
	}

	hashedPassword, err := encryption.HashPassword(req.Password, salt)

	if err != nil {
		return entity.GeneralErrorWithAdditionalInfo(err.Error())
	}

	if err = s.repo.UpdatePassword(hashedPassword, salt, userID); err != nil {
		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}
	
	return nil
}