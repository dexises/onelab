package repository

import "onelab/internal/model"

type IUserRepository interface {
	Create(user *model.User) error
	Get(id int) (*model.User, error)
}

type IbookRepository interface{}

type Manager struct {
	User IUserRepository
	Book IbookRepository
}

func NewRepository() *Manager {
	return &Manager{
		User: NewUserRepository(),
		// Book: NewBookRepository(),
	}
}
