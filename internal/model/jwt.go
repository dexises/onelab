package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
