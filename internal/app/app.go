package app

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunApp(port string) {
	//cfg := config.Config{}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)

	eventRepo := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepo)

	contactRepo := repository.NewContactRepository()
	contactService := service.NewContactService(contactRepo)
	// Ручки
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create", eventService.CreateEvent)
	e.GET("/events", eventService.GetAllEvents)

	e.POST("/contacts", contactService.CreateContact)
	e.GET("/contacts", contactService.GetAllContacts)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
