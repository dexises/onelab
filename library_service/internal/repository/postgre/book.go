package postgre

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"onelab/internal/model"
)

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{
		DB: db,
	}
}

func (r *BookRepo) Get(ctx context.Context, id int) (model.Book, error) {
	var book model.Book
	if err := r.DB.Table("books").Unscoped().First(&book, id).Error; err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, model.ErrRecordNotFound
		}
		return model.Book{}, err
	}
	return book, nil
}

func (r *BookRepo) Create(ctx context.Context, book model.Book) (uint, error) {
	result := r.DB.WithContext(ctx).Table("books").Create(&book)
	if result.Error != nil {
		return 0, result.Error
	}
	return book.ID, nil
}

func (r *BookRepo) GetAll(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	if err := r.DB.Table("books").Find(&books).Error; err != nil {
		return []model.Book{}, err
	}
	return books, nil
}
