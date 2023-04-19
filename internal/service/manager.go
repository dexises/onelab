package service

import (
	"onelab/internal/config"
	"onelab/internal/repository"
)

type Manager struct {
	User    IUserService
	Book    IBookService
	Library ILibraryService
}

func NewService(cfg *config.Config, repo *repository.Manager) *Manager {
	return &Manager{
		User:    NewUserService(repo),
		Book:    NewBookService(repo),
		Library: NewLibraryService(repo),
	}
}
