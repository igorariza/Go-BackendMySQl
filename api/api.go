package api

import (
	"github.com/igorariza/Go-BackendMySQl/internal/storage"
	users "github.com/igorariza/Go-BackendMySQl/users/web"
)

func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()

	r := routes(users.NewUserHTTPService(db))
	server := newServer(port, r)

	server.Start()
}
