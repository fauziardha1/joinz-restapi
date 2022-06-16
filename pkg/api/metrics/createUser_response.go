package metrics

import "joinz-api/pkg/db"

// CreateUserResponse is a struct that represents a response from the create user endpoint
type CreateUserResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	User    *db.User `json:"user"`
}

// CreateUserRequest is a struct that represents a request to create a new user
type CreateUserRequest struct {
	Email    string `json:"email"`
	Uname    string `json:"uname"`
	Password string `json:"password"`
	Phone    string `json:"phone,omitempty"`
}
