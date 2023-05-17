package services

import (
	"database/sql"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/ilhamnyto/echo-fw/entity"
	"github.com/ilhamnyto/echo-fw/repositories"
)

type InterfacePostService interface {
	CreatePost(req *entity.CreatePostRequest, userId int) *entity.CustomError
	GetAllPost(cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError)
	GetPost(postId string) (*entity.PostData, *entity.CustomError)
	GetUserPost(username string, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError)
	GetUserPostByUsernameOrBody(query string, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError)
	GetMyPost(userId int, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError)// GetMyPost(userId int, cursor string) ([]*entity.UserPost, *entity.CustomError)
}

type PostService struct {
	repo repositories.InterfacePostRepository
}

func NewPostService(repo repositories.InterfacePostRepository) InterfacePostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(req *entity.CreatePostRequest, userId int) *entity.CustomError {

	post := entity.Post{
		UserID: userId,
		Body: req.Body,
		CreatedAt: time.Now(),

	}

	if err := s.repo.Create(&post); err != nil {
		return entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (s *PostService) GetAllPost(cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError) {
	var (
		t time.Time
		userPosts []*entity.UserPost
		err error
	)

	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		
		if err != nil {
			return nil, nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
		}
		timestamp := int64(cursorInt)
		t = time.Unix(timestamp, 0).UTC()
		userPosts, err = s.repo.GetAllPost(&t)
	}else{
		userPosts, err = s.repo.GetAllPost(nil)
	}


	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, entity.NotFoundError()
		}

		return nil, nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var postsData []*entity.PostData
	if len(userPosts) == 6 {
		for _, post := range userPosts[:5] {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}else {
		for _, post := range userPosts {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}


	paging := entity.Paging{}

	if len(userPosts) == 6 {
		paging.Next = true
		paging.Cursor = strconv.Itoa(int(postsData[len(postsData)-1].CreatedAt.Unix()))
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return postsData, &paging, nil
}

func (s *PostService) GetPost(postId string) (*entity.PostData, *entity.CustomError) {
	decodedPostId, err := base64.StdEncoding.DecodeString(postId)

	if err != nil {
		return nil, entity.BadRequestErrorWithAdditionalInfo(err.Error())
	}
	
	id, err := strconv.Atoi(strings.Split(string(decodedPostId), ":")[1])
	
	if err != nil {
		return nil, entity.BadRequestErrorWithAdditionalInfo(err.Error())
	}

	post, err := s.repo.GetPost(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundError()
		}

		return nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	postData := entity.PostData{}
	postData.ParseEntityToResponse(post)

	return &postData, nil
}

func (s *PostService) GetUserPost(username string, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError) {
	var (
		t time.Time
		userPosts []*entity.UserPost
		err error
	)

	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		
		if err != nil {
			return nil, nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
		}
		timestamp := int64(cursorInt)
		t = time.Unix(timestamp, 0).UTC()
		userPosts, err = s.repo.GetUserPost(username, &t)
	}else{
		userPosts, err = s.repo.GetUserPost(username, nil)
	}


	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, entity.NotFoundError()
		}

		return nil, nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var postsData []*entity.PostData
	if len(userPosts) == 6 {
		for _, post := range userPosts[:5] {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}else {
		for _, post := range userPosts {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}


	paging := entity.Paging{}

	if len(userPosts) == 6 {
		paging.Next = true
		paging.Cursor = strconv.Itoa(int(postsData[len(postsData)-1].CreatedAt.Unix()))
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return postsData, &paging, nil
}

func (s *PostService) GetUserPostByUsernameOrBody(query string, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError) {
	var (
		t time.Time
		userPosts []*entity.UserPost
		err error
	)

	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		
		if err != nil {
			return nil, nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
		}
		timestamp := int64(cursorInt)
		t = time.Unix(timestamp, 0).UTC()
		userPosts, err = s.repo.GetUserPostByUsernameOrBody(query, &t)
	}else{
		userPosts, err = s.repo.GetUserPostByUsernameOrBody(query, nil)
	}


	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, entity.NotFoundError()
		}

		return nil, nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var postsData []*entity.PostData
	if len(userPosts) == 6 {
		for _, post := range userPosts[:5] {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}else {
		for _, post := range userPosts {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}


	paging := entity.Paging{}

	if len(userPosts) == 6 {
		paging.Next = true
		paging.Cursor = strconv.Itoa(int(postsData[len(postsData)-1].CreatedAt.Unix()))
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return postsData, &paging, nil
}

func (s *PostService) GetMyPost(userId int, cursor string) ([]*entity.PostData, *entity.Paging, *entity.CustomError) {
	var (
		t time.Time
		userPosts []*entity.UserPost
		err error
	)

	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		
		if err != nil {
			return nil, nil, entity.GeneralErrorWithAdditionalInfo(err.Error())
		}
	
		timestamp := int64(cursorInt)
		t = time.Unix(timestamp, 0).UTC()
		userPosts, err = s.repo.GetUserPostByUserId(userId, &t)
	}else{
		userPosts, err = s.repo.GetUserPostByUserId(userId, nil)
	}


	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, entity.NotFoundError()
		}

		return nil, nil, entity.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var postsData []*entity.PostData
	if len(userPosts) == 6 {
		for _, post := range userPosts[:5] {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}else {
		for _, post := range userPosts {
			tempPost := new(entity.PostData)
			tempPost.ParseEntityToResponse(post)
			postsData = append(postsData, tempPost)		
		}

	}


	paging := entity.Paging{}

	if len(userPosts) == 6 {
		paging.Next = true
		paging.Cursor = strconv.Itoa(int(postsData[len(postsData)-1].CreatedAt.Unix()))
	}else {
		paging.Next = false
		paging.Cursor = ""
	}
	return postsData, &paging, nil
}