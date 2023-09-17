package helpers

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}