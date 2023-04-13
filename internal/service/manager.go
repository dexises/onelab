package service

import (
	"onelab/internal/repository"
)

type Manager struct {
	User    IUserService
	Book    IBookService
	Library ILibraryService
}

func NewService(repo *repository.Manager) *Manager {
	return &Manager{
		User:    NewUserService(repo),
		Book:    NewBookService(repo),
		Library: NewLibraryService(repo),
	}
}
