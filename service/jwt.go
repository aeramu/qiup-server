package service

import(
	"github.com/dgrijalva/jwt-go"
)
  
var JWTSecretKey = []byte("KentangGorengB3raKs1")
  
type JWTClaims struct{
	jwt.StandardClaims
	Payload string
}

func GenerateJWT(payload string)(string){
	jwtClaims := &JWTClaims{
		Payload: payload,
	}
	token,_ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString(JWTSecretKey)
	return "token="+token
}  

func DecodeJWT(token string)(string){
	token = token[6:]
	claims := new(JWTClaims)
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	return claims.Payload
}