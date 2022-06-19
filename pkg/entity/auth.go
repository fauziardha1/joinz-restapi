package entity

import "joinz-api/pkg/repository/db"

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	User    *db.User `json:"user"`
}
