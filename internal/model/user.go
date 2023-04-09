package model

import (
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`
}
