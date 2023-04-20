package model

import "time"

type TransactionCreate struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	BookID      uint      `json:"book_id"`
	BookPrice   uint      `json:"book_price"`
	UserBalance uint      `json:"user_balance"`
	Date        time.Time `json:"transfer_date" gorm:"not null"`
}
type TransactionResponse struct {
	ID      uint   `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
