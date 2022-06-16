package metrics

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
