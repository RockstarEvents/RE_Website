package service

import (
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	event := new(models.Event)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := s.repo.CreateEvent(userID, *event); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, event)
}

// GetAllEvents обрабатывает получение всех событий
func (s *eventService) GetAllEvents(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	events, err := s.repo.GetAllEvents(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}