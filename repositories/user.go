package repositories

import (
	"database/sql"
	// "fmt"
	"time"

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

	queryGetUserCredentialByUsername = `
		SELECT id, username, password, salt from users where username = $1
	`
	queryGetUserByUsername = `
		SELECT username, COALESCE(first_name,''), COALESCE(last_name,''), 
		email, COALESCE(phone_number,''), COALESCE(location,''), date_trunc('second',created_at) from users
		where username = $1
	`
)


type InterfaceUserRepository interface {
	Create(user entity.User) (error)
	CheckUsernameAndEmail(username string, email string) (int, error)
	GetUserCredentialByUsername(username string) (*entity.User, error)
	GetAllUser(cursor *time.Time) ([]*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
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

func (r *UserRepository) GetUserCredentialByUsername(username string) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserCredentialByUsername)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	user := entity.User{}

	if err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Salt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUser(cursor *time.Time) ([]*entity.User, error) {
	queryGetAllUser := `
		SELECT username, COALESCE(first_name,''), COALESCE(last_name,''), email, COALESCE(phone_number,''), COALESCE(location,''), date_trunc('second',created_at) from users
	`

	if cursor != nil {
		queryGetAllUser += " where created_at < '" + cursor.Format("2006-01-02 15:04:05") +"'"
	}

	queryGetAllUser += " order by created_at desc limit 6"


	stmt, err := r.db.Prepare(queryGetAllUser)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		tempUser := new(entity.User)
		
		err := rows.Scan(
							&tempUser.Username, &tempUser.FirstName, &tempUser.LastName, 
							&tempUser.Email, &tempUser.PhoneNumber, &tempUser.Location, &tempUser.CreatedAt,
						)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUser)
	}

	return users, nil

}

func (r *UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserByUsername)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	user := entity.User{}

	if err := row.Scan(
		&user.Username, &user.FirstName, &user.LastName, 
		&user.Email, &user.PhoneNumber, &user.Location, &user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}