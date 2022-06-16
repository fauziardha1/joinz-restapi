package db

import "github.com/go-pg/pg/v10"

// User is a struct that represents a user
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Uname     string `json:"uname"`
	Password  string `json:"password"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateUser(db *pg.DB, req *User) (*User, error) {

	// insert new user to db
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	// get the new user from db
	// Select user by primary key.
	user := &User{ID: req.ID}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}
