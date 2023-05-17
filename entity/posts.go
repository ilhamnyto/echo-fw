package entity

import (
	"encoding/base64"
	"time"
)

type Post struct {
	UserID    int    		`json:"post_id"`
	Body      string 		`json:"body"`
	CreatedAt time.Time		`json:"created_at"`
	DeletedAt time.Time		`json:"deleted_at"`
}

type UserPost struct {
	PostID 		string		`json:"post_id"`
	Username	string		`json:"username"`
	Body 		string		`json:"body"`
	CreatedAt	time.Time	`json:"created_at"`
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

func (p *PostData) ParseEntityToResponse(u *UserPost) {
	p.PostID = base64.StdEncoding.EncodeToString([]byte(u.Username+":"+u.PostID))
	p.Username = u.Username
	p.Body = u.Body
	p.CreatedAt = u.CreatedAt
}