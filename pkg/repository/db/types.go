package db

// Contact is a struct that represents a contact
type Contact struct {
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	User            *User  `pg:"rel:has-one" json:"user"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Relationship    string `json:"relationship"`
	ProfilePicture  string `json:"profile_picture"`
	Priorities      int    `json:"priorities"`
	AlternativeName string `json:"alternative_name"`
	Messages        string `json:"messages"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// User is a struct that represents a user
type User struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	Uname          string `json:"uname"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
	FullName       string `json:"full_name"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Nickname       string `json:"nickname"`
	UserAddress    string `json:"user_address"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
