package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	BookID     uint
	Borrower   string
	DueDate    time.Time
	ReturnDate *time.Time
}
