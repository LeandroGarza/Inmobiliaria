package dto

import "github.com/golang-jwt/jwt"

type Claims struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}
