package model

import (
	"errors"
	"time"
)

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	CreatedAt    time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt    time.Time `json:"updated_at" swaggerignore:"true"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"not null;unique" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
