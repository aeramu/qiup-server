package service

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("KentangGorengB3raKs1")

type jwtClaims struct {
	jwt.StandardClaims
	Payload string
}

//GenerateJWT token
func GenerateJWT(payload string) string {
	jwtClaims := &jwtClaims{
		Payload: payload,
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString(jwtSecretKey)
	return "token=" + token
}

//DecodeJWT token
func DecodeJWT(token string) string {
	token = token[6:]
	claims := new(jwtClaims)
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	return claims.Payload
}
