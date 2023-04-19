package postgre

import (
	"context"
	"errors"
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

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) Create(ctx context.Context, user model.User) (uint, error) {
	result := r.DB.WithContext(ctx).Table("users").Omit("updated_at").Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *UserRepo) Get(ctx context.Context, id int) (model.User, error) {
	var user model.User
	if err := r.DB.Table("users").Unscoped().Omit("password_hash").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, model.ErrRecordNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user model.Login) error {
	result := r.DB.WithContext(ctx).Table("users").Where("email = ?", user.Email).Update("password_hash", user.Password)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return model.ErrRecordNotFound
	}
	return nil
}

func (r *UserRepo) TestCreate(ctx context.Context, user model.User) (string, error) {
	result := r.DB.WithContext(ctx).Table("users").Omit("deleted_at").Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Email, nil
}
