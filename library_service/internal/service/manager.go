package service

import (
	"onelab/internal/config"
	"onelab/internal/repository"
)

type Manager struct {
	User        IUserService
	Book        IBookService
	Transaction ITransactionService
	Library     ILibraryService
}

func NewService(cfg *config.Config, repo *repository.Manager) *Manager {
	user := NewUserService(repo)
	book := NewBookService(repo)
	transaction := NewTransactionService(cfg)
	library := NewLibraryService(repo, transaction, book, user)
	return &Manager{
		User:        user,
		Book:        book,
		Transaction: transaction,
		Library:     library,
	}
}
