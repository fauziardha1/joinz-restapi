package metrics

type CreateUserResponse struct {
	Status    string `json:"status"`
	Success   bool   `json:"success"`
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Uname     string `json:"uname"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
