package repository

import "onelab/internal/model"

type userRepository struct {
	db map[int]model.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: map[int]model.User{},
	}
}

func (r *userRepository) Create() {
}

func (r *userRepository) Get() {
}

func (r *userRepository) Update() {
}

func (r *userRepository) Delete() {
}
