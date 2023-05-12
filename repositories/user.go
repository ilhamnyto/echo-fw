package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/echo-fw/entity"
)

var (
	queryCreateUser = `
		INSERT INTO users (username, email, password, salt, created_at) 
		VALUES ($1, $2, $3, $4, $5)
	`

	queryCheckUsernameAndEmail = `
		SELECT COUNT(username) from users where username = $1 or email = $2
	`
)


type InterfaceUserRepository interface {
	Create(user entity.User) (error)
	CheckUsernameAndEmail(username string, email string) (int, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) InterfaceUserRepository {
	return &UserRepository{db: db}
}


func (r *UserRepository) Create(user entity.User) (error) {
	stmt, err := r.db.Prepare(queryCreateUser)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(user.Username, user.Email, user.Password, user.Salt, user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CheckUsernameAndEmail(username string, email string) (int, error) {
	stmt, err := r.db.Prepare(queryCheckUsernameAndEmail)

	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(username, email)

	var exist int

	if err = row.Scan(&exist); err != nil {
		return 0, err
	}

	return exist, nil
}

