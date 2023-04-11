package postgre

import "gorm.io/gorm"

type BookRepo struct {
	DB *gorm.DB
}
