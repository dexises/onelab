package service

import (
	"onelab/internal/config"
	"onelab/internal/repository"
)

type Manager struct {
	User    IUserService
	Book    IBookService
	Library ILibraryService
	Jwt     IAuthService
}

func NewService(cfg *config.Config, repo *repository.Manager) *Manager {
	return &Manager{
		Jwt:     NewJWT(cfg),
		User:    NewUserService(repo),
		Book:    NewBookService(repo),
		Library: NewLibraryService(repo),
	}
}
