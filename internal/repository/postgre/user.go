package postgre

import (
	"context"

	"onelab/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: DB,
	}
}

func (r *UserRepo) Create(ctx context.Context, user model.User) error {
	return r.DB.WithContext(ctx).Omit("deleted_at").Create(&user).Error
}

func (r *UserRepo) Get(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user model.User) error {
	result := r.DB.WithContext(ctx).Where("email = ?", user.Email).Update("password_hash", user.PasswordHash)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
