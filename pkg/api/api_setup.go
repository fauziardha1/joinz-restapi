package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupApi() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// set routes
	r.Route("/users", func(r chi.Router) {

		r.Post("/", createArticle) // POST /users

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

// func createNewUser
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create new user"))
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
