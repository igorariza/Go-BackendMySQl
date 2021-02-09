package middlew

import (
	"net/http"

	st "github.com/igorariza/Go-BackendMySQl/internal/storage"
)

//ChequeoBD es el middlew y permite conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := st.ConnectToDB().Ping()
		if err != nil {
			http.Error(w, "Conexión perdida con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)

	}
}
