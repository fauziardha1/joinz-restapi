package users

import (
	"encoding/json"
	"fmt"
	"joinz-api/pkg/entity"
	"joinz-api/pkg/repository/db"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
)

// func createNewUser is a function that creates a new user
// it takes a request body of type CreateUserRequest
// it returns a success when the user is created, otherwise it returns an error
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// parse request body
	req := &entity.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &entity.CreateUserResponse{
			Success: false,
			Error:   err.Error(),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &entity.CreateUserResponse{
			Success: false,
			Error:   err.Error(),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error get database from context: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get time
	t := time.Now()

	// insert new user to db
	user, err := db.CreateUser(pgdb, &db.User{
		Email:     req.Email,
		Uname:     req.Uname,
		Password:  req.Password,
		Phone:     req.Phone,
		CreatedAt: fmt.Sprintf("%s", t.Format("20060102150405")),
		UpdatedAt: fmt.Sprintf("%s", t.Format("20060102150405")),
	})
	if err != nil {
		res := &entity.CreateUserResponse{
			Success: false,
			Error:   err.Error(),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error insert data to db: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return response

	res := &entity.CreateUserResponse{
		Success: true,
		Error:   "",
		User:    user,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	return
}

// func GetAllUsers is a function that gets all users
// it takes no parameters
// it returns a success when the users are retrieved, otherwise it returns an error
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &entity.CreateUserResponse{
			Success: false,
			Error:   "failed to get database from context",
			User:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error get database from context: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get all users from db
	users, err := db.GetAllUsers(pgdb)
	if err != nil {
		res := &entity.CreateUserResponse{
			Success: false,
			Error:   err.Error(),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error insert data to db: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := &entity.GetAllUsersResponse{
		Success: true,
		Error:   "",
		Users:   users,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	return
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all user"))
}

// func getUserByEmail is a function that gets a user by email
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	w.Write([]byte(fmt.Sprintf("get user by email: %s", email)))
}

// func getUserByUname is a function that gets a user by uname
func GetUserByUname(w http.ResponseWriter, r *http.Request) {
	uname := chi.URLParam(r, "uname")
	w.Write([]byte(fmt.Sprintf("get user by uname: %s", uname)))
}

// func getUserById is a function that gets a user by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("get user by id: %s", id)))
}

// func updateUser
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("update user by id: %s", id)))
}

// func deleteUser is a function that deletes a user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("delete user by id: %s", id)))
}
