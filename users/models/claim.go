package users

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim es la estrucutura usada para procesar el jwt
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `json:"idUser"`
	jwt.StandardClaims
}
