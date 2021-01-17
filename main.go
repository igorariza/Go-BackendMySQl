package main

import (
	"log"

	"github.com/igorariza/Go-BackendMySQl/bd"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}

}
