package service

import (
	
	"eventPlanner/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactService interface {
	GetAllContacts(c echo.Context) error
	SelectContacts(c echo.Context) error
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) GetAllContacts(c echo.Context) error {
	contacts, err := s.repo.GetAllContacts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"contacts": contacts})
}

func (s *contactService) SelectContacts(c echo.Context) error {
	var request struct {
		ContactIDs []string `json:"contact_ids"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	contacts, err := s.repo.SelectContacts(request.ContactIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":          true,
		"selected_contacts": contacts,
	})
}
