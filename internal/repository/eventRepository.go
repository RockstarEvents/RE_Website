package repository

import (
	"errors"
	"eventPlanner/internal/models"
	"sync"
)

// EventRepository интерфейс для репозитория событий
type EventRepository interface {
	CreateEvent(userID int64, event models.Event) error
	GetAllEvents(userID int64) ([]models.Event, error)
}

// InMemoryEventRepository реализация EventRepository в памяти
type InMemoryEventRepository struct {
	mu     sync.RWMutex
	events map[int64][]models.Event
	nextID int
}

// NewEventRepository создает новый InMemoryEventRepository
func NewEventRepository() EventRepository {
	return &InMemoryEventRepository{
		events: make(map[int64][]models.Event),
		nextID: 1,
	}
}

// CreateEvent добавляет событие в репозиторий для указанного пользователя
func (r *InMemoryEventRepository) CreateEvent(userID int64, event models.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	event.ID = r.nextID
	r.nextID++
	r.events[userID] = append(r.events[userID], event)
	return nil
}

// GetAllEvents возвращает все события для указанного пользователя
func (r *InMemoryEventRepository) GetAllEvents(userID int64) ([]models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	events, ok := r.events[userID]
	if !ok || len(events) == 0 {
		return nil, errors.New("no events found for this user")
	}

	return events, nil
}
