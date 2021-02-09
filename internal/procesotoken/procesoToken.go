package internal

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	usr "github.com/igorariza/Go-BackendMySQl/users/models"
)

//Email Comment generic
var Email string

//IDUsuario Comment generic
var IDUsuario string

//ProcesoToken procesa token para obtener sus valores
func ProcesoToken(tk string) (*usr.Claim, bool, string, error) {
	miClave := []byte("Groupar_instagram")
	claims := &usr.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
		return claims, true, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
