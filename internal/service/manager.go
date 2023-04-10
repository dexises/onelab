package service

import (
	"onelab/internal/repository"
)

type Manager struct {
	User IUserService
	Book IbookService
}

func NewService(repo *repository.Manager) *Manager {
	return &Manager{
		User: NewUserService(repo),
	}
}
