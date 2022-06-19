package db

type Contact struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	User            *User  `pg:"rel:has-one" json:"user"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Relationship    string `json:"relationship"`
	ProfilePicture  string `json:"profile_picture"`
	Priorities      string `json:"priorities"`
	AlternativeName string `json:"alternative_name"`
	Messages        string `json:"messages"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}