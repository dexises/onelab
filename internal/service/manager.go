package service

import (
	"context"

	"onelab/internal/model"
	"onelab/internal/repository"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User) error
}

type IbookRepository interface{}

type Manager struct {
	User IUserService
	Book IbookRepository
}

func NewService(repo *repository.Manager) (*Manager, error) {
	return &Manager{
		User: NewUserService(repo.User),
	}
}
