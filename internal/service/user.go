package service

import (
	"context"

	"onelab/internal/jsonlog"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository, logger *jsonlog.Logger) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user model.User) error {
	// validation

	err := s.repo.Create(ctx, user)
	if err != nil {
		s.logger.PrintError(err, map[string]string{
			"Name":  user.Name,
			"Login": user.Login,
		})
		return err
	}

	s.logger.PrintInfo("user saved to database", map[string]string{
		"Name":  user.Name,
		"Login": user.Login,
	})
	return nil
}

func (s *UserService) Get(id int) (*model.User, error) {
	return s.repo.Get(id)
}
