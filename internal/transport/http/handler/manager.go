package handler

import (
	"onelab/internal/config"
	"onelab/internal/service"
	"onelab/internal/transport/http/middleware"
)

type Manager struct {
	srv     *service.Manager
	JWTAuth *middleware.JWTAuth
}

func NewManager(cfg *config.Config, srv *service.Manager) *Manager {
	return &Manager{
		srv:     srv,
		JWTAuth: middleware.NewJWTAuth(cfg, srv.User),
	}
}
