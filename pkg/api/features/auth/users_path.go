package auth

import (
	"encoding/json"
	"fmt"
	"joinz-api/pkg/api/metrics"
	"joinz-api/pkg/db"
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

// func createNewUser is a function that creates a new user
// it takes a request body of type CreateUserRequest
// it returns a success when the user is created, otherwise it returns an error
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("create new user"))

	// parse request body
	req := &metrics.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &metrics.CreateUserResponse{
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
		res := &metrics.CreateUserResponse{
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
		res := &metrics.CreateUserResponse{
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

	res := &metrics.CreateUserResponse{
		Success: true,
		Error:   "",
		User:    user,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	return
}
