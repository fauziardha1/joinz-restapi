package db

type Contact struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	User         *User  `pg:"rel:has-one" json:"user"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Relationship string `json:"relationship"`
	Message      string `json:"message"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
