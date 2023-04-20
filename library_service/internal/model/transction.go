package model

type TransactionCreate struct {
	UserID      uint `json:"user_id"`
	BookID      uint `json:"book_id"`
	BookPrice   uint `json:"book_price"`
	UserBalance uint `json:"user_balance"`
}

type TransactionResponse struct {
	ID      uint   `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
