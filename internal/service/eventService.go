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

// CreateEvent обрабатывает создание нового события
func (s *eventService) CreateEvent(c echo.Context) error {
	event := new(models.Event)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := s.repo.CreateEvent(*event); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "event created")
}

// GetAllEvents обрабатывает получение всех событий
func (s *eventService) GetAllEvents(c echo.Context) error {
	events, err := s.repo.GetAllEvents()
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, events)
}
