package postgre

import (
	"context"
	"gorm.io/gorm"
	"onelab/internal/model"
	"time"
)

type LibraryRepo struct {
	DB *gorm.DB
}

func NewLibraryRepo(db *gorm.DB) *LibraryRepo {
	return &LibraryRepo{
		DB: db,
	}
}

func (r *LibraryRepo) GiveBook(ctx context.Context, item model.CreateBookReader) error {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Получить пользователя
	var user model.User
	if err := tx.Table("users").Where("id = ?", item.ReaderID).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Получить книгу
	var book model.Book
	if err := tx.Table("books").Where("id = ?", item.BookID).First(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Обновить баланс пользователя и доход книги
	user.Balance -= book.Price
	book.Income += book.Price

	// Создать запись в books_readers
	item.TakenAt = time.Now()
	if err := tx.Table("books_readers").Create(&item).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Обновить пользователя и книгу
	if err := tx.Table("users").Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table("books").Where("id = ?", book.ID).Updates(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Сохранить изменения в базе данных
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *LibraryRepo) ReturnBook(ctx context.Context, bookReader model.BookReader) error {
	result := r.DB.WithContext(ctx).Table("books_readers").Where("reader_id = ? AND book_id = ?", bookReader.ReaderID, bookReader.BookID).Update("returned_at", bookReader.ReturnedAt)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *LibraryRepo) List(ctx context.Context) ([]model.BookRentSummary, error) {
	var bookRentSummaries []model.BookRentSummary
	if err := r.DB.Table("books_readers").
		Select("books.title, books_readers.reader_id, books.income").
		Joins("JOIN books ON books.id = books_readers.book_id").
		Where("books_readers.returned_at IS NULL").
		Scan(&bookRentSummaries).Error; err != nil {
		return nil, err
	}

	return bookRentSummaries, nil
}
