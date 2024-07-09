package Implementations

import (
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) service.EventService {
	return &eventService{
		repo: repo,
	}
}

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
