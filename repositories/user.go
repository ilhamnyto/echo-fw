package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/echo-fw/entity"
)

type InterfaceUserRepository interface {
	Create(user entity.User) (error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) InterfaceUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user entity.User) (error) {


	return nil
}

