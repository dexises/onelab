package service

import (
	"context"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type IBookService interface {
	Create(ctx context.Context, book model.Book) error
	Get(ctx context.Context, id int) (model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
}

type BookService struct {
	Repo *repository.Manager
}

func NewBookService(repo *repository.Manager) *BookService {
	return &BookService{
		Repo: repo,
	}
}

func (s BookService) Create(ctx context.Context, book model.Book) error {
	//TODO validation
	err := s.Repo.Book.Create(ctx, book)
	return err
}

func (s BookService) Get(ctx context.Context, id int) (model.Book, error) {
	book, err := s.Repo.Book.Get(ctx, id)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (s BookService) GetAll(ctx context.Context) ([]model.Book, error) {
	books, err := s.Repo.Book.GetAll(ctx)
	if err != nil {
		return []model.Book{}, err
	}
	return books, nil
}
