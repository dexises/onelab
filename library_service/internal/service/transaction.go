package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"onelab/internal/config"
	"onelab/internal/model"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, bookID uint, userID uint, price uint, balance uint) (uint, error)
}

type TransactionService struct {
	client *http.Client
	url    string
}

func NewTransactionService(cfg *config.Config) *TransactionService {
	return &TransactionService{
		client: &http.Client{},
		url:    cfg.TransactionServiceURL,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, bookID uint, userID uint, price uint, balance uint) (uint, error) {
	data := &model.TransactionCreate{
		UserID:      userID,
		BookID:      bookID,
		BookPrice:   price,
		UserBalance: balance,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url+"/transactions", bytes.NewReader(jsonData))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var transactionResponse model.TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionResponse)
	if err != nil {
		return 0, err
	}

	if !transactionResponse.Success {
		return 0, errors.New(fmt.Sprintf("Failed to create transaction: %s", transactionResponse.Message))
	}

	return transactionResponse.ID, nil
}
