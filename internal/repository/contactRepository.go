package repository

import (
	"errors"
	"eventPlanner/internal/models"
	"sync"
)

// ContactRepository интерфейс для репозитория контактов
type ContactRepository interface {
	CreateContact(contact models.Contact) error
	GetAllContacts() ([]models.Contact, error)
}

// InMemoryContactRepository реализация ContactRepository в памяти
type InMemoryContactRepository struct {
	mu       sync.RWMutex
	contacts []models.Contact
}

// NewContactRepository создает новый InMemoryContactRepository
func NewContactRepository() ContactRepository {
	return &InMemoryContactRepository{
		contacts: []models.Contact{},
	}
}

// CreateContact добавляет контакт в репозиторий
func (r *InMemoryContactRepository) CreateContact(contact models.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, c := range r.contacts {
		if c.Email == contact.Email {
			return errors.New("contact with this email already exists")
		}
	}

	r.contacts = append(r.contacts, contact)
	return nil
}

// GetAllContacts возвращает все контакты
func (r *InMemoryContactRepository) GetAllContacts() ([]models.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.contacts) == 0 {
		return nil, errors.New("no contacts found")
	}

	return r.contacts, nil
}
