package service

import (
	"onelab/internal/config"
	"onelab/internal/repository"
)

type Manager struct {
	Transactions ITransactionService
}

func NewService(cfg *config.Config, repo *repository.Manager) *Manager {
	return &Manager{
		Transactions: NewTransactionService(repo),
	}
}
