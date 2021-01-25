package storage

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	users "github.com/igorariza/Go-BackendMySQl/users/models"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t users.User) (string, error) {

	miClave := []byte(t.Email)

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.FirstName,
		"apellidos":        t.LastName,
		"fecha_nacimiento": t.DateBirth,
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
