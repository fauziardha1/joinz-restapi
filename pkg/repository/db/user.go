package db

import "github.com/go-pg/pg/v10"

// func CreateUser is a function that creates a new user
// it takes a request body of type CreateUserRequest
// it returns a success when the user is created, otherwise it returns an error
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

// func GetUserByID is a function that gets all users
// it takes no parameters
// it returns a list of users when success, otherwise it returns an error
func GetAllUsers(db *pg.DB) ([]*User, error) {
	var users []*User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// func Login is a function that logs in a user
// it takes a request body of type LoginRequest
// it returns a success when the user is logged in, otherwise it returns an error
func GetUserByEmailOrUserName(db *pg.DB, username string, email string) (*User, error) {

	// get the user from db where username or email is equal to the username or email\
	if username == "" {
		username = email
	} else {
		email = username
	}
	user := &User{}
	err := db.Model(user).Where("uname = ?", username).WhereOr("email = ?", email).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}
