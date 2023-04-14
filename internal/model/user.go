package model

import (
	"errors"
	"time"
)

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null" json:"-"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
