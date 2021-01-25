package api

import (
	"github.com/go-chi/chi"
	"github.com/igorariza/Go-BackendMySQl/internal/storage"
	users "github.com/igorariza/Go-BackendMySQl/users/web"
)

//Start comment generic
func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()

	r := chi.NewMux()

	routesCreateUser(users.NewUserCreateHTTPService(db), r)
	routesLoginUser(users.NewUserLoginHTTPService(db), r)
	server := newServer(port, r)

	server.Start()
}
