package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedpassword string, plainpassword string) bool {
	fmt.Printf("hashedpassword: %v\n", hashedpassword)
	fmt.Printf("plainpassword: %v\n", plainpassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(plainpassword))
	if err != nil {
		fmt.Printf("bcrypt error: %v", err)
		return false
	}
	return true
}

type Claims struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userid int, username string) (string, error) {
	claims := &Claims{
		Userid:   userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// creacion de token con claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// clave secreta
	jwtkey := []byte("tengohambre")

	// token jwt firmado con la clave
	tokenstring, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
