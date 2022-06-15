package metrics

type UpdateUserResponse struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
}
