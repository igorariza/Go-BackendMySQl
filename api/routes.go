package api

import (
	"github.com/go-chi/chi"
	users "github.com/igorariza/Go-BackendMySQl/users/web"
)

func routesCreateUser(services *users.UserCreateHTTPService, r *chi.Mux) {

	// r.Get("/users", services.GetUsersHandler)
	r.Post("/users", services.CreateUsersHandler)
	r.Get("/users/{userID}", services.GetUsersByIDHandler)

}

func routesLoginUser(services *users.LoginUserHTTPService, r *chi.Mux) {

	r.Post("/login", services.LoginUsersHandler)
	// r.Post("/users", services.CreateUsersHandler)
	// r.Get("/users/{userID}", services.GetUsersByIDHandler)

}
