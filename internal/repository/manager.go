package repository

type IUserRepository interface {
	Create()
	Delete()
	Update()
	Get()
}

type Repository struct {
	User IUserRepository
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}
