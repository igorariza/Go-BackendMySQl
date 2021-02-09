package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/igorariza/Go-BackendMySQl/internal/jwt"
	"github.com/igorariza/Go-BackendMySQl/internal/web"
	users "github.com/igorariza/Go-BackendMySQl/users/gateway"
	usr "github.com/igorariza/Go-BackendMySQl/users/models"
)

//LoginUserHTTPService comment generic
type LoginUserHTTPService struct {
	gtw users.LoginGateway
}

//NewUserLoginHTTPService comment generic
func NewUserLoginHTTPService(db *sql.DB) *LoginUserHTTPService {
	return &LoginUserHTTPService{
		users.NewLoginUserGateway(db),
	}
}

//LoginUsersHandler comment generic
func (s *LoginUserHTTPService) LoginUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var cmd usr.LoginUser
	body := r.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(cmd.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	user, err := s.gtw.LoginUser(&cmd)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos ", 400)
		return
	}
	if user == nil {
		http.Error(w, "Usuario y/o Contraseña inválidos ", http.StatusBadRequest)
		return
	}
	jwtKey, err := jwt.GeneroJWT(*user)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}
	resp := usr.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	//Grabacion de cookie en el lado Usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
	web.Success(&user, http.StatusOK).Send(w)

}
