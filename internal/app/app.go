package app

import (
	"context"
	"eventPlanner/internal/config"
	"eventPlanner/internal/database"
	"eventPlanner/internal/repository/Implementations"
	Implementations2 "eventPlanner/internal/service/Implementations"
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

	dsn := cfg.DataBasePath
	conn, err := database.ConnectDB(dsn)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	userRepo := Implementations.NewUserRepository(conn)
	eventRepo := Implementations.NewEventRepository(conn)
	contactRepo := Implementations.NewContactRepository(conn) // Используем новый репозиторий контактов

	userService := Implementations2.NewUserService(userRepo)
	eventService := Implementations2.NewEventService(eventRepo)
	contactService := Implementations2.NewContactService(contactRepo) // Инициализируем сервис контактов с новым репозиторием

	// Ручки
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create/:id", eventService.CreateEvent)
	e.GET("/events/:id", eventService.GetAllEvents)

	e.GET("/contacts", contactService.GetAllContacts)         // Добавляем ручку для получения всех контактов
	e.POST("/contacts/select", contactService.SelectContacts) // Добавляем ручку для выбора контактов

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
