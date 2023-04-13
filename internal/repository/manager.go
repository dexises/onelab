package repository

import (
	"context"

	"onelab/internal/config"
	"onelab/internal/model"
	"onelab/internal/repository/postgre"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.User) error
}

type IBookRepository interface {
	Create(ctx context.Context, book model.Book) error
	Get(ctx context.Context, id int) (model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
}

type ILibraryRepository interface {
	GiveBook(ctx context.Context, item model.CreateBookReader) error
	ReturnBook(ctx context.Context, bookReader model.BookReader) error
	List(ctx context.Context) ([]model.BookReader, error)
	ListMonth(ctx context.Context) ([]model.BookReader, error)
}

type Manager struct {
	pg      *gorm.DB
	User    IUserRepository
	Book    IBookRepository
	Library ILibraryRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Manager, error) {
	db, err := postgre.NewConnection(cfg.DB)
	if err != nil {
		return nil, err
	}

	return &Manager{
		pg:      db,
		User:    postgre.NewUserRepo(db),
		Book:    postgre.NewBookRepo(db),
		Library: postgre.NewLibraryRepo(db),
	}, nil
}
