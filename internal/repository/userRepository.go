package repository

import (
	"eventPlanner/internal/models"
	"fmt"
	"sync"
)

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	FindUser(username string) (*models.User, error)
}

// хранение в памяти
type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  []models.User
	nextID int
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) CreateUser(user models.User) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Name == user.Name {
			return 0, fmt.Errorf("username already exists")
		}
	}
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user.ID, nil
}

func (r *InMemoryUserRepository) FindUser(username string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Name == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
