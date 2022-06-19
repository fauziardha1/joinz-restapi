package entity

import "joinz-api/pkg/repository/db"

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

// DeleteUserResponse is a struct that represents a response from the delete user endpoint
type DeleteUserResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
}

// GetUserResponse is a struct that represents a response from the get user endpoint
type GetUserResponse struct {
	Status    string `json:"status"`
	Success   bool   `json:"success"`
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Uname     string `json:"uname"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllUsersResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	Users   []*db.User `json:"users"`
}

// UpdateUserResponse is a struct that represents a response from the update user endpoint
type UpdateUserResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Email   string `json:"email,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Phone   string `json:"phone,omitempty"`
}
