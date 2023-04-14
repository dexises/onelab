package service

import (
	"github.com/golang-jwt/jwt/v4"
	"onelab/internal/config"
	"onelab/internal/model"
	"time"
)

type IAuthService interface {
	GenerateToken(username string) (string, error)
}

type JWT struct {
	jwtKey []byte
}

func NewJWT(cfg *config.Config) *JWT {
	return &JWT{
		jwtKey: []byte(cfg.JwtToken),
	}
}

func (m *JWT) GenerateToken(email string) (tokenStr string, err error) {
	expireTime := 24 * time.Hour
	claims := &model.JwtCustomClaims{
		Name: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.jwtKey)
}
