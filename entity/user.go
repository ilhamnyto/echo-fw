package entity

import "time"

type User struct {
	ID          int    		`json:"id"`
	Username    string 		`json:"username"`
	Email       string 		`json:"email"`
	Password 	string 		`json:"password"`
	Salt 		string		`json:"salt"`
	FirstName   string 		`json:"first_name"`
	LastName    string 		`json:"last_name"`
	PhoneNumber string 		`json:"phone_number"`
	Location    string 		`json:"location"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type CreateUserRequest struct {
	Username	string		`json:"username"`
	Email		string 		`json:"email"`
	Password	string		`json:"password"`
}


type UpdateUserRequest struct {
	FirstName 	string		`json:"first_name"`
	LastName    string 		`json:"last_name"`
	PhoneNumber string 		`json:"phone_number"`
	Location    string 		`json:"location"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type UpdatePasswordRequest struct {
	Password			string		`json:"password"`
	ConfirmPassword		string 		`json:"confirm_password"`
}

type Paging struct {
	Next	bool	`json:"next"`
	Cursor	int		`json:"cursor"`
}

type UserData struct {
	Username    string 		`json:"username"`
	Email       string 		`json:"email"`
	FirstName   string 		`json:"first_name"`
	LastName    string 		`json:"last_name"`
	PhoneNumber string 		`json:"phone_number"`
	Location    string 		`json:"location"`
	CreatedAt   time.Time	`json:"created_at"`
}

type UserResponse struct {
	Data 	interface{}		`json:"data"`
	Paging	Paging			`json:"paging"`	
}

