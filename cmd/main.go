package main

import (
	"log"
	"os"

	"github.com/igorariza/Go-BackendMySQl/api"
	"github.com/igorariza/Go-BackendMySQl/config"
)

const defaultPort = "8080"

func main() {
	config.LoadConfig()
	log.Println("starting API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
