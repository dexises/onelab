package service

import (
	"onelab/internal/jsonlog"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type UserService struct {
	repo   repository.IUserRepository
	logger *jsonlog.Logger
}

func NewUserService(repo repository.IUserRepository, logger *jsonlog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) Create(user *model.User) error {
	s.logger.PrintInfo("Creating user", map[string]string{
		"Name":  user.Name,
		"Login": user.Login,
	})

	//validation

	s.logger.PrintInfo("saving user to database", map[string]string{
		"Name":  user.Name,
		"Login": user.Login,
	})
	err := s.repo.Create(user)
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
