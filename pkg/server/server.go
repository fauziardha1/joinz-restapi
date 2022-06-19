package server

import (
	"joinz-api/pkg/server/httphandler/auth"
	"joinz-api/pkg/server/httphandler/users"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
)

func SetupApi(pgDB *pg.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgDB))

	return defineRoutes(r)
}

// func defineRoute is a function that defines a route
func defineRoutes(r *chi.Mux) *chi.Mux {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", users.CreateNewUser) // POST /users

		r.Get("/", users.GetAllUsers)                   // GET /users/
		r.Get("/byemail/{email}", users.GetUserByEmail) // GET /users/byemail/:email
		r.Get("/byuname/{uname}", users.GetUserByUname) // GET /users/byid/:uname
		r.Get("/byid/{id}", users.GetUserByID)          // GET /users/byid/:id

		r.Put("/{id}", users.UpdateUserByID) // PUT /users/:id (update user by id)
		r.Delete("/{id}", users.DeleteUser)  // DELETE /users/:id

	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", users.CreateNewUser) // POST /auth/register
		r.Post("/login", auth.Login)             // POST /auth/login
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}
