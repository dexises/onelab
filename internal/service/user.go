package service

import (
	"onelab/internal/model"
	"onelab/internal/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) Get(id int) (*model.User, error) {
	return s.repo.Get(id)
}
