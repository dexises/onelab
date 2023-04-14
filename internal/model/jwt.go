package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
