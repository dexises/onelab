package repository

import (
	"sync"

	"onelab/internal/model"
)

type UserRepository struct {
	mu      sync.RWMutex
	counter int
	db      map[int]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: make(map[int]*model.User),
	}
}

func (r *UserRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	user.ID = r.counter
	r.db[user.ID] = user

	return nil
}

func (r *UserRepository) Get(id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.db[id]
	if !ok {
		return nil, model.ErrRecordNotFound
	}

	return user, nil
}
