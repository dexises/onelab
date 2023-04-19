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
	item.TakenAt = time.Now()
	return r.DB.WithContext(ctx).Table("books_readers").Create(&item).Error
}

func (r *LibraryRepo) ReturnBook(ctx context.Context, bookReader model.BookReader) error {
	result := r.DB.WithContext(ctx).Table("books_readers").Where("reader_id = ?", bookReader.ReaderID).Update("returned_at", bookReader.ReturnedAt)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *LibraryRepo) List(ctx context.Context) ([]model.BookReader, error) {
	var allReaders []model.BookReader
	if err := r.DB.WithContext(ctx).Table("books_readers").Find(&allReaders).Error; err != nil {
		return []model.BookReader{}, err
	}
	return allReaders, nil
}

func (r *LibraryRepo) ListMonth(ctx context.Context) ([]model.BookReader, error) {
	var readers []model.BookReader
	if err := r.DB.WithContext(ctx).Table("books_readers").Select("readers.id, readers.name, COUNT(*) as books_taken_last_month").
		Joins("LEFT JOIN readers ON readers.id = book_readers.reader_id").
		Where("book_readers.taken_at BETWEEN ? AND ?", time.Now().AddDate(0, -1, 0), time.Now()).
		Group("readers.id").
		Order("books_taken_last_month DESC").Find(&readers).Error; err != nil {
		return nil, err
	}

	return readers, nil
}
