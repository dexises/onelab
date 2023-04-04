package service

import (
	"onelab/internal/jsonlog"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type IUserService interface {
	Create(user *model.User) error
	Get(id int) (*model.User, error)
}

type Service struct {
	User IUserService
}

func NewService(repo *repository.Manager, logger *jsonlog.Logger) *Service {
	return &Service{
		User: NewUserService(repo.User, logger),
	}
}
