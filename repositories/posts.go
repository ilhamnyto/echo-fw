package repositories

import (
	"database/sql"
	"time"

	"github.com/ilhamnyto/echo-fw/entity"
)

var (
	queryCreatePost = `
		INSERT INTO posts (user_id, body, created_at) VALUES ($1, $2, $3)
	`
	queryGetPostById = `
		SELECT p.id, u.username, p.body, date_trunc('second',p.created_at) from posts as p LEFT JOIN users as u ON p.user_id = u.id where p.id = $1
	`
)

type InterfacePostRepository interface {
	Create(post *entity.Post) error
	GetAllPost(cursor *time.Time) ([]*entity.UserPost, error)
	GetPost(postId int) (*entity.UserPost, error)
	GetUserPost(username string, cursor *time.Time) ([]*entity.UserPost, error)
	GetUserPostByUsernameOrBody(query string, cursor *time.Time) ([]*entity.UserPost, error)
	GetUserPostByUserId(userId int, cursor *time.Time) ([]*entity.UserPost, error)
}

type PostRepository struct {
	db	*sql.DB
}

func NewPostRepository(db *sql.DB) InterfacePostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *entity.Post) error {
	stmt, err := r.db.Prepare(queryCreatePost)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(post.UserID, post.Body, post.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) GetAllPost(cursor *time.Time) ([]*entity.UserPost, error) {
	queryGetAllPost := `
		SELECT p.id, u.username, p.body, date_trunc('second',p.created_at) from posts as p LEFT JOIN users as u ON p.user_id = u.id
	`

	if cursor != nil {
		queryGetAllPost += " where p.created_at < '" + cursor.Format("2006-01-02 15:04:05") +"'"
	}

	queryGetAllPost += " order by p.created_at desc limit 6"


	stmt, err := r.db.Prepare(queryGetAllPost)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)
		if err := rows.Scan(&tempPost.PostID, &tempPost.Username, &tempPost.Body, &tempPost.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, tempPost)
	}

	return posts, nil
}

func (r *PostRepository) GetPost(postId int) (*entity.UserPost, error) {
	stmt, err := r.db.Prepare(queryGetPostById)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(postId)

	post := entity.UserPost{}

	if err = row.Scan(&post.PostID, &post.Username, &post.Body, &post.CreatedAt); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) GetUserPost(username string, cursor *time.Time) ([]*entity.UserPost, error) {
	queryGetUserPost := `
		SELECT p.id, u.username, p.body, p.created_at from posts as p LEFT JOIN users as u ON p.user_id = u.id where
	`

	if cursor != nil {
		queryGetUserPost += " p.created_at < '" + cursor.Format("2006-01-02 15:04:05") +"' and"
	}

	queryGetUserPost += " u.username = $1 order by p.created_at desc limit 6"

	stmt, err := r.db.Prepare(queryGetUserPost)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)
		if err := rows.Scan(&tempPost.PostID, &tempPost.Username, &tempPost.Body, &tempPost.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, tempPost)
	}

	return posts, nil
}

func (r *PostRepository) GetUserPostByUsernameOrBody(query string, cursor *time.Time) ([]*entity.UserPost, error) {
	query = "%" + query + "%" 

	queryGetUserPostByUsernameOrBody := `
		SELECT p.id, u.username, p.body, p.created_at from posts as p LEFT JOIN users as u ON p.user_id = u.id where
	`

	if cursor != nil {
		queryGetUserPostByUsernameOrBody += " p.created_at < '" + cursor.Format("2006-01-02 15:04:05") +"' and"
	}

	queryGetUserPostByUsernameOrBody += " (u.username like $1 or p.body like $1) order by p.created_at desc limit 6"


	stmt, err := r.db.Prepare(queryGetUserPostByUsernameOrBody)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)
		if err := rows.Scan(&tempPost.PostID, &tempPost.Username, &tempPost.Body, &tempPost.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, tempPost)
	}

	return posts, nil
}

func (r *PostRepository) GetUserPostByUserId(userId int, cursor *time.Time) ([]*entity.UserPost, error) {
	queryGetUserPost := `
		SELECT p.id, u.username, p.body, p.created_at from posts as p LEFT JOIN users as u ON p.user_id = u.id where
	`

	if cursor != nil {
		queryGetUserPost += " p.created_at < '" + cursor.Format("2006-01-02 15:04:05") +"' and"
	}

	queryGetUserPost += " p.user_id = $1 order by p.created_at desc limit 6"

	stmt, err := r.db.Prepare(queryGetUserPost)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.UserPost

	for rows.Next() {
		tempPost := new(entity.UserPost)
		if err := rows.Scan(&tempPost.PostID, &tempPost.Username, &tempPost.Body, &tempPost.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, tempPost)
	}

	return posts, nil
}