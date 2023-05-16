package repositories

import (
	"database/sql"
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
	queryGetUserByUserID = `
		SELECT username, COALESCE(first_name,''), COALESCE(last_name,''), 
		email, COALESCE(phone_number,''), COALESCE(location,''), date_trunc('second',created_at) from users
		where id = $1
	`
	queryUpdateUser = `
		UPDATE users SET first_name = $1, last_name = $2, phone_number = $3, location = $4, updated_at = $5 where id = $6
	`
	queryUpdatePassword = `
		UPDATE users SET password = $1, salt = $2, updated_at = $3 where id = $4
	`
	queryGetUserPasswordAndSalt = `
		SELECT password, salt from users where id = $1
	`
)


type InterfaceUserRepository interface {
	Create(user entity.User) (error)
	CheckUsernameAndEmail(username string, email string) (int, error)
	GetUserCredentialByUsername(username string) (*entity.User, error)
	GetAllUser(cursor *time.Time) ([]*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	GetUserBySimalarUsernameOrEmail(query string, cursor *time.Time) ([]*entity.User, error)
	GetUserByUserID(userID int) (*entity.User, error)
	UpdateUser(req *entity.UpdateUserRequest, userId int) error
	UpdatePassword(password string, salt string, userId int) error
	GetUserPasswordAndSalt(userID int) (*entity.User, error)
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

func (r *UserRepository) GetUserBySimalarUsernameOrEmail(query string, cursor *time.Time) ([]*entity.User, error) {
	queryGetUserBySimilarUsernameOrEmail := `
		SELECT username, COALESCE(first_name,''), COALESCE(last_name,''), email, COALESCE(phone_number,''), COALESCE(location,''), date_trunc('second',created_at) from users where
	`
	query = "%" + query + "%"
	
	if cursor != nil {
		queryGetUserBySimilarUsernameOrEmail += " created_at < '" + cursor.Format("2006-01-02 15:04:05") +"' and"
	}

	queryGetUserBySimilarUsernameOrEmail += " (username like $1 or email like $1) order by created_at desc limit 6"	

	stmt, err := r.db.Prepare(queryGetUserBySimilarUsernameOrEmail)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(query)

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

func (r *UserRepository) GetUserByUserID(userID int) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserByUserID)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userID)

	user := entity.User{}

	err = row.Scan(
		&user.Username, &user.FirstName, &user.LastName, 
		&user.Email, &user.PhoneNumber, &user.Location, &user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *UserRepository) UpdateUser(req *entity.UpdateUserRequest, userId int) error {
	stmt, err := r.db.Prepare(queryUpdateUser)

	if err != nil {
		return err
	}
	
	updatedAt := time.Now()
	if _, err := stmt.Exec(req.FirstName, req.LastName, req.PhoneNumber, req.Location, updatedAt, userId); err != nil {
		return err
	}	

	return nil
}

func (r *UserRepository) UpdatePassword(password string, salt string, userId int) error {
	stmt, err := r.db.Prepare(queryUpdatePassword)

	if err != nil {
		return err
	}

	updatedAt := time.Now()
	if _, err := stmt.Exec(password, salt, updatedAt, userId); err != nil {
		return err
	}
	
	return nil
}

func (r *UserRepository) GetUserPasswordAndSalt(userID int) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserPasswordAndSalt)
	
	if err != nil {
		return nil, err
	}
	
	row := stmt.QueryRow(userID)

	user := entity.User{}

	if err := row.Scan(&user.Password, &user.Salt); err != nil {
		return nil, err
	}

	return &user, nil
}