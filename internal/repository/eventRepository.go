package repository

import (
	"eventPlanner/internal/models"
	"sync"
)

// EventRepository интерфейс для репозитория событий
type EventRepository interface {
	CreateEvent(event models.Event) error
	GetAllEvents() ([]models.Event, error)
}

// InMemoryEventRepository реализация EventRepository в памяти
type InMemoryEventRepository struct {
	mu     sync.RWMutex
	events []models.Event
	nextID int
}

// NewEventRepository создает новый InMemoryEventRepository
func NewEventRepository() EventRepository {
	return &InMemoryEventRepository{
		events: []models.Event{},
		nextID: 1,
	}
}

// CreateEvent добавляет событие в репозиторий
func (r *InMemoryEventRepository) CreateEvent(event models.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	event.ID = r.nextID
	r.nextID++
	r.events = append(r.events, event)
	return nil
}

// GetAllEvents возвращает все события
func (r *InMemoryEventRepository) GetAllEvents() ([]models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.events, nil
}
