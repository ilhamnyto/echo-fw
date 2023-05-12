package entity

import "time"

type Post struct {
	UserID    int    		`json:"post_id"`
	Body      string 		`json:"body"`
	CreatedAt time.Time		`json:"created_at"`
	DeletedAt time.Time		`json:"deleted_at"`
}

type CreatePostRequest struct {
	UserID		int 		`json:"user_id"`
	Body		string		`json:"body"`
	CreatedAt	time.Time	`json:"created_at"`
}

type PostData struct {
	PostID		string		`json:"post_id"`
	Username	string 		`json:"username"`
	Body 		string		`json:"body"`
	CreatedAt	time.Time	`json:"created_at"`
}