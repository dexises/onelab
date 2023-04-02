package service

import (
	"fmt"
	"onelab/internal/repository"
)

type IUserService interface {
	Create()
	Delete()
	Update()
	Get()
}

type Service struct {
	User IUserService
}

func NewService(repo repository.Repository) *Service {
	fmt.Println(repo.User.)
	return &Service{
		User: repo.User,
	}
}
