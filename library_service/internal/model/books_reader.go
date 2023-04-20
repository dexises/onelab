package model

import "time"

type CreateBookReader struct {
	BookID        uint      `json:"book_id" gorm:"column:book_id"`
	ReaderID      uint      `json:"reader_id" gorm:"column:reader_id"`
	TransactionID uint      `swaggerignore:"true"`
	TakenAt       time.Time ` swaggerignore:"true"`
	CreatedAt     time.Time ` swaggerignore:"true"`
	UpdatedAt     time.Time ` swaggerignore:"true"`
}

type BookReader struct {
	ID            uint      `swaggerignore:"true"`
	BookID        uint      `json:"book_id" gorm:"column:book_id"`
	ReaderID      uint      `json:"reader_id" gorm:"column:reader_id"`
	TransactionID uint      `swaggerignore:"true"`
	BookTitle     string    `json:"book_title" gorm:"column:book_title"`
	TakenAt       time.Time `swaggerignore:"true"`
	ReturnedAt    time.Time `swaggerignore:"true"`
	CreatedAt     time.Time `swaggerignore:"true"`
	UpdatedAt     time.Time `swaggerignore:"true"`
}
