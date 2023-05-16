package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/echo-fw/entity"
)

var (
	queryCreatePost = `
		INSERT INTO posts (user_id, body, created_at) VALUES ($1, $2, $3)
	`

	queryGetAllPost = `
		SELECT 
	`
)

type InterfacePostRepository interface {
	Create(post *entity.Post) error
	GetAllPost() ([]*entity.Post, error)
	GetPost(postId int) (*entity.Post, error)
	GetUserPost(username string) ([]*entity.Post, error)
	GetUserPostByUsernameOrBody(query string) ([]*entity.Post, error)
}

type PostRepository struct {
	db	*sql.DB
}

func NewPostRepository(db *sql.DB) InterfacePostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *entity.Post) error {
	return nil
}

func (r *PostRepository) GetAllPost() ([]*entity.Post, error) {
	return nil, nil
}

func (r *PostRepository) GetPost(postId int) (*entity.Post, error) {
	return nil, nil
}

func (r *PostRepository) GetUserPost(username string) ([]*entity.Post, error) {
	return nil, nil
}

func (r *PostRepository) GetUserPostByUsernameOrBody(query string) ([]*entity.Post, error) {
	return nil, nil
}