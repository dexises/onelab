package repository

import (
	"context"

	"onelab/internal/config"
	"onelab/internal/model"
	"onelab/internal/repository/postgre"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transactionsCreate model.TransactionCreate) (uint, error)
	GetAll(ctx context.Context) ([]model.TransactionCreate, error)
}

type Manager struct {
	pg           *gorm.DB
	Transactions ITransactionRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Manager, error) {
	db, err := postgre.NewConnection(cfg.DB)
	if err != nil {
		return nil, err
	}

	return &Manager{
		pg:           db,
		Transactions: postgre.NewTransactionRepo(db),
	}, nil
}
