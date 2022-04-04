package middleware

import (
	"project/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userid int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userid
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
