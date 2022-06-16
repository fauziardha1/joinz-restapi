package api

import (
	"encoding/json"
	"fmt"
	"joinz-api/pkg/api/metrics"
	"joinz-api/pkg/db"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
)

func SetupApi(pgDB *pg.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgDB))

	// set routes
	r.Route("/users", func(r chi.Router) {

		r.Post("/", createNewUser) // POST /users

		r.Get("/", getAllUser)                    // GET /users/
		r.Get("/byemail/{email}", getUserByEmail) // GET /users/byemail/:email
		r.Get("/byuname/{uname}", getUserByUname) // GET /users/byid/:uname
		r.Get("/byid/{id}", getUserByID)          // GET /users/byid/:id

		r.Put("/{id}", updateUserByID) // PUT /users/:id (update user by id)
		r.Delete("/{id}", deleteUser)  // DELETE /users/:id

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all user"))
}

// func getUserByEmail is a function that gets a user by email
func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	w.Write([]byte(fmt.Sprintf("get user by email: %s", email)))
}

// func getUserByUname is a function that gets a user by uname
func getUserByUname(w http.ResponseWriter, r *http.Request) {
	uname := chi.URLParam(r, "uname")
	w.Write([]byte(fmt.Sprintf("get user by uname: %s", uname)))
}

// func getUserById is a function that gets a user by id
func getUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("get user by id: %s", id)))
}

// func createNewUser is a function that creates a new user
// it takes a request body of type CreateUserRequest
// it returns a success when the user is created, otherwise it returns an error
func createNewUser(w http.ResponseWriter, r *http.Request) {
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

// func updateUser
func updateUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("update user by id: %s", id)))
}

// func deleteUser is a function that deletes a user by id
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("delete user by id: %s", id)))
}
