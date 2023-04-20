package service

import (
	"context"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, transactionsCreate model.TransactionCreate) (uint, error)
	GetAll(ctx context.Context) ([]model.TransactionCreate, error)
}

type TransactionService struct {
	Repo *repository.Manager
}

func NewTransactionService(repo *repository.Manager) *TransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (s TransactionService) CreateTransaction(ctx context.Context, transactions model.TransactionCreate) (uint, error) {
	//TODO validation
	id, err := s.Repo.Transactions.CreateTransaction(ctx, transactions)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s TransactionService) GetAll(ctx context.Context) ([]model.TransactionCreate, error) {
	trn, err := s.Repo.Transactions.GetAll(ctx)
	if err != nil {
		return []model.TransactionCreate{}, err
	}
	return trn, nil
}
