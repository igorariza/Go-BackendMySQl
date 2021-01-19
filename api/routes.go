package api

import (
	"github.com/go-chi/chi"
	users "github.com/igorariza/Go-BackendMySQl/users/web"
)

func routes(services *users.UserHTTPService) *chi.Mux {
	r := chi.NewMux()

	//r.Get("/users", services.GetUsersHandler)
	//r.Post("/users", services.CreateUsersHandler)
	// r.Get("/users/{userID}", services.GetUsersByIDHandler)
	r.Get("/users/{email}", services.GetUsersByEmailHandler)

	return r
}
