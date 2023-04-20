package postgre

import (
	"context"
	"onelab/internal/model"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: db,
	}
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, transactions model.TransactionCreate) (uint, error) {
	result := r.DB.WithContext(ctx).Table("transactions").Create(&transactions)
	if result.Error != nil {
		return 0, result.Error
	}
	return transactions.ID, nil
}

func (r *TransactionRepo) GetAll(ctx context.Context) ([]model.TransactionCreate, error) {
	var transactions []model.TransactionCreate
	if err := r.DB.Table("transactions").Find(&transactions).Error; err != nil {
		return []model.TransactionCreate{}, err
	}
	return transactions, nil
}
