package server

import (
	"joinz-api/pkg/server/httphandler/auth"
	"joinz-api/pkg/server/httphandler/contacts"
	"joinz-api/pkg/server/httphandler/users"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
)

func InitApp(pgDB *pg.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgDB))

	return defineRoutes(r)
}

// func defineRoute is a function that defines a route
func defineRoutes(r *chi.Mux) *chi.Mux {

	// define route for users endpoint
	r.Route("/users", func(r chi.Router) {
		// Post Routes
		r.Post("/", users.CreateNewUser) // POST /users

		// Get Routes
		r.Get("/", users.GetAllUsers) // GET /users/
		// TODO: implement handler for endpoints below
		r.Get("/byemail/{email}", users.GetUserByEmail) // GET /users/byemail/:email
		r.Get("/byuname/{uname}", users.GetUserByUname) // GET /users/byid/:uname
		r.Get("/byid/{id}", users.GetUserByID)          // GET /users/byid/:id

		// TODO: implement handler for endpoints below
		// Put Routes
		r.Put("/{id}", users.UpdateUserByID) // PUT /users/:id (update user by id)

		// TODO: implement handler for endpoints below
		// Delete Routes
		r.Delete("/{id}", users.DeleteUser) // DELETE /users/:id

	})

	// define route for auth endpoint
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", users.CreateNewUser) // POST /auth/register
		r.Post("/login", auth.Login)             // POST /auth/login
	})

	// define route for contacts endpoint
	r.Route("/contacts", func(r chi.Router) { // GET /contacts/
		r.Get("/{userid}", contacts.GetContactsByUserID) // GET /contacts/byuserid/:userid
		r.Post("/", contacts.CreateNewContact)           // POST /contacts
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to JOINZ API!!\nIt's creates by JOINZ team:\n\t-Devon\n\t-fauziarda\n\t-Heri\n\t-Rafi\n\t-Ricky \nPlease use /docs to see the available endpoints\nhttps://github.com/fauziardha1/joinz-restapi"))
	})

	return r
}
