package app

import (
	"context"
	"eventPlanner/internal/config"
	"eventPlanner/internal/database"
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"log"
)

func RunApp(cfg config.Config) {

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dsn := "postgres://postgres:postgres@localhost:5432/postgres"
	conn, err := database.ConnectDB(dsn)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	userRepo := repository.NewUserRepository(conn)
	eventRepo := repository.NewEventRepository(conn)
	contactRepo := repository.NewContactRepository(conn) // Используем новый репозиторий контактов

	userService := service.NewUserService(userRepo)
	eventService := service.NewEventService(eventRepo)
	contactService := service.NewContactService(contactRepo) // Инициализируем сервис контактов с новым репозиторием

	// Ручки
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create", eventService.CreateEvent)
	e.GET("/events", eventService.GetAllEvents)

	e.GET("/contacts", contactService.GetAllContacts)         // Добавляем ручку для получения всех контактов
	e.POST("/contacts/select", contactService.SelectContacts) // Добавляем ручку для выбора контактов

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
