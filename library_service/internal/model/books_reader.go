package model

import "time"

type CreateBookReader struct {
	BookID        uint `json:"book_id" gorm:"column:book_id"`
	ReaderID      uint `json:"reader_id" gorm:"column:reader_id"`
	TransactionID uint
	TakenAt       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type BookReader struct {
	ID            uint
	BookID        uint `json:"book_id" gorm:"column:book_id"`
	ReaderID      uint `json:"reader_id" gorm:"column:reader_id"`
	TransactionID uint
	BookTitle     string    `json:"book_title" gorm:"column:book_title"`
	TakenAt       time.Time //`json:"taken_at"`
	ReturnedAt    time.Time //`json:"returned_at"`
	CreatedAt     time.Time //`json:"created_at"`
	UpdatedAt     time.Time
}
