package auth

import (
	"encoding/json"
	"fmt"
	"joinz-api/pkg/entity"
	"joinz-api/pkg/repository/db"
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
)

// Login is a function that logs in a user
func Login(w http.ResponseWriter, r *http.Request) {
	// parse request body
	req := &entity.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &entity.LoginResponse{
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

	// check if username and email is empty
	if req.Email == "" && req.UserName == "" {
		res := &entity.LoginResponse{
			Success: false,
			Error:   "email or username is not allowed to be empty",
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

	// get the user by email or username
	user, err := db.GetUserByEmailOrUserName(pgdb, req.Email, req.UserName)
	if err != nil {
		res := &entity.LoginResponse{
			Success: false,
			Error:   fmt.Sprintf("error getting user by email or username: %s, %s, %s", err, req.Email, req.UserName),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// check if user is nil
	if user == nil {
		res := &entity.LoginResponse{
			Success: false,
			Error:   "username or password is incorrect (user not found)",
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if password is correct
	if user.Password != req.Password {
		res := &entity.LoginResponse{
			Success: false,
			Error:   "username or password is incorrect",
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error encoding response: %s\n", err)
		}

		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// return the user
	res := &entity.LoginResponse{
		Success: true,
		Error:   "",
		User:    user,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	return
}
