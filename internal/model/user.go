package model

import "errors"

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	ID       int
	Name     string
	Login    string
	Password []byte
}
