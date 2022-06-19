package entity

// DeleteUserResponse is a struct that represents a response from the delete user endpoint
type DeleteUserResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
}
