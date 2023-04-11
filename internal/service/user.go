package service

import (
	"context"
	"errors"

	"onelab/internal/model"
	"onelab/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User, id int) error
}

type UserService struct {
	repo *repository.Manager
}

func NewUserService(repo *repository.Manager) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user model.User) error {
	// TODO user validation
	var err error
	user.PasswordHash, err = s.SetPassword(user.PasswordHash)
	if err != nil {
		return err
	}

	return s.repo.User.Create(ctx, user)
}

func (s *UserService) Get(ctx context.Context, id int) (model.User, error) {
	if id < 1 {
		return model.User{}, errors.New("invalid id parameter")
	}

	user, err := s.repo.User.Get(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, user model.User, id int) error {
	extUser, err := s.repo.User.Get(ctx, id)
	if err != nil {
		return err
	}

	extUser.PasswordHash, err = s.SetPassword(user.PasswordHash)
	if err != nil {
		return errors.New("generate password error")
	}

	if extUser.Email != user.Email {
		return errors.New("Email address is not correct")
	}

	return s.repo.User.Update(ctx, extUser, id)
}

func (s *UserService) SetPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
