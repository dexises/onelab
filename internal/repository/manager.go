package repository

import "onelab/internal/model"

type IUserRepository interface {
	Create(user *model.User) error
	Get(id int) (*model.User, error)
}

type Manager struct {
	User IUserRepository
}

func NewRepository() *Manager {
	return &Manager{
		User: NewUserRepository(),
	}
}
