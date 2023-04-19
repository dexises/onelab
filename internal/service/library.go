package service

import (
	"context"
	"onelab/internal/model"
	"onelab/internal/repository"
	"time"
)

type ILibraryService interface {
	GiveBook(ctx context.Context, item model.CreateBookReader) error
	ReturnBook(ctx context.Context, bookReader model.BookReader) error
	List(ctx context.Context) ([]model.BookReader, error)
	ListMonth(ctx context.Context) ([]model.BookReader, error)
}

type LibraryService struct {
	Repo *repository.Manager
}

func NewLibraryService(repo *repository.Manager) *LibraryService {
	return &LibraryService{
		Repo: repo,
	}
}

func (s LibraryService) GiveBook(ctx context.Context, item model.CreateBookReader) error {
	//validation
	return s.Repo.Library.GiveBook(ctx, item)
}

func (s LibraryService) ReturnBook(ctx context.Context, bookReader model.BookReader) error {
	//validation
	bookReader.ReturnedAt = time.Now()
	if err := s.Repo.Library.ReturnBook(ctx, bookReader); err != nil {
		return err
	}
	return nil
}

func (s LibraryService) List(ctx context.Context) ([]model.BookReader, error) {
	bookReaders, err := s.Repo.Library.List(ctx)
	if err != nil {
		return []model.BookReader{}, err
	}
	return bookReaders, nil
}

func (s LibraryService) ListMonth(ctx context.Context) ([]model.BookReader, error) {
	bookReadersMonth, err := s.Repo.Library.ListMonth(ctx)
	if err != nil {
		return []model.BookReader{}, err
	}
	return bookReadersMonth, nil
}
