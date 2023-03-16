package jwt

import "github.com/dgrijalva/jwt-go/v4"

/*
Набор полей для токена. jwt.StandardClaims описаны все стандартные поля согласно спецификации JWT.
*/
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
