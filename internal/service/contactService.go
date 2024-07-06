package service

import (
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ContactService интерфейс для сервиса контактов
type ContactService interface {
	CreateContact(c echo.Context) error
	GetAllContacts(c echo.Context) error
}

// contactService реализация ContactService
type contactService struct {
	repo repository.ContactRepository
}

// NewContactService создает новый contactService
func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{
		repo: repo,
	}
}

// CreateContact обрабатывает создание нового контакта
func (s *contactService) CreateContact(c echo.Context) error {
	contact := new(models.Contact)
	if err := c.Bind(contact); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if err := s.repo.CreateContact(*contact); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, contact)
}

// GetAllContacts обрабатывает получение всех контактов
func (s *contactService) GetAllContacts(c echo.Context) error {
	contacts, err := s.repo.GetAllContacts()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, contacts)
}
