package models

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
