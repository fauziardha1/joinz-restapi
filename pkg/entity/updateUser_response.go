package entity

// UpdateUserResponse is a struct that represents a response from the update user endpoint
type UpdateUserResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Email   string `json:"email,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Phone   string `json:"phone,omitempty"`
}
