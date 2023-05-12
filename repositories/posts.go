package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/echo-fw/entity"
)

type InterfacePostRepository interface {
	Create(post *entity.Post) error
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