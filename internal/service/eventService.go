package service

import (
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

// EventService интерфейс для сервиса событий
type EventService interface {
	CreateEvent(c echo.Context) error
	GetAllEvents(c echo.Context) error
}

// eventService реализация EventService
type eventService struct {
	repo repository.EventRepository
}

// NewEventService создает новый eventService
func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{
		repo: repo,
	}
}

// CreateEvent создает новое событие
func (s *eventService) CreateEvent(c echo.Context) error {
	e := new(models.Event)
	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := s.repo.CreateEvent(*e); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	// TODO: функция,которая попросит участников мероприятия из сервиса Contacts
	// TODO: после получения списка участников проходимся по ним и добавляем создаём это мероприятие у них
	return c.JSON(http.StatusCreated, "created")
}

// GetAllEvents возвращает все события
func (s *eventService) GetAllEvents(c echo.Context) error {
	events, err := s.repo.GetAllEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}
