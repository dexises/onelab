package service

import (
	"context"
	"errors"
	"fmt"
	"onelab/internal/model"
	"onelab/internal/repository"
	"time"
)

type ILibraryService interface {
	GiveBook(ctx context.Context, item model.CreateBookReader) error
	ReturnBook(ctx context.Context, bookReader model.BookReader) error
	List(ctx context.Context) ([]model.BookRentSummary, error)
}

type LibraryService struct {
	UserService        *UserService
	BookService        *BookService
	TransactionService *TransactionService
	Repo               *repository.Manager
}

func NewLibraryService(repo *repository.Manager, transactionService *TransactionService, bookService *BookService, UserService *UserService) *LibraryService {
	return &LibraryService{
		UserService:        UserService,
		BookService:        bookService,
		TransactionService: transactionService,
		Repo:               repo,
	}
}

func (s LibraryService) GiveBook(ctx context.Context, item model.CreateBookReader) error {
	//validation
	var book model.Book
	book, err := s.BookService.Get(ctx, int(item.BookID))
	if err != nil {
		return errors.New(fmt.Sprint("book does not exist"))
	}

	var user model.User
	user, err = s.UserService.Get(ctx, int(item.ReaderID))
	if err != nil {
		return errors.New(fmt.Sprint("user does not exist"))
	}

	if book.Price > user.Balance {
		return errors.New(fmt.Sprint("Not enough balance to borrow book"))
	}
	//когда делаешь запрос на http://localhost:9090/transactions с одного микросервиса на другой он все время возвращает
	//ошибку connect: connection refused, однако когда делаешь запрос с постмана все работает, не могу понять с чем ошибка
	//item.TransactionID, err = s.TransactionService.CreateTransaction(ctx, item.BookID, item.ReaderID, book.Price, user.Balance)
	//if err != nil {
	//	return err
	//}
	return s.Repo.Library.GiveBook(ctx, item)
}

func (s LibraryService) ReturnBook(ctx context.Context, bookReader model.BookReader) error {
	//validation
	book, err := s.BookService.Get(ctx, int(bookReader.BookID))
	if err != nil {
		return errors.New(fmt.Sprint("book does not exist"))
	}

	if book.Title != bookReader.BookTitle {
		return errors.New(fmt.Sprint("wrong book title"))
	}

	_, err = s.UserService.Get(ctx, int(bookReader.ReaderID))
	if err != nil {
		return errors.New(fmt.Sprint("user does not exist"))
	}

	bookReader.ReturnedAt = time.Now()
	if err := s.Repo.Library.ReturnBook(ctx, bookReader); err != nil {
		return err
	}
	return nil
}

func (s LibraryService) List(ctx context.Context) ([]model.BookRentSummary, error) {
	bookReaders, err := s.Repo.Library.List(ctx)
	if err != nil {
		return []model.BookRentSummary{}, err
	}
	return bookReaders, nil
}
