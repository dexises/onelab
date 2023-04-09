package repository

import (
	"context"

	"onelab/internal/config"
	"onelab/internal/model"
	"onelab/internal/repository/postgre"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User) error
}

type IbookRepository interface{}

type Manager struct {
	pg   *gorm.DB
	User IUserRepository
	Book IbookRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Manager, error) {
	db, err := postgre.NewConnection(cfg.DB)
	if err != nil {
		return nil, err
	}

	return &Manager{
		pg:   db,
		User: postgre.NewUserRepo(db),
	}, nil
}
