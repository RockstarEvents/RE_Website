package app

import (
	"context"
	"eventPlanner/internal/Router"
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

	dsn := cfg.DataBasePath
	conn, err := database.ConnectDB(dsn)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	userRepo := Implementations.NewUserRepository(conn)
	eventRepo := Implementations.NewEventRepository(conn)
	contactRepo := Implementations.NewContactRepository(conn)

	userService := Implementations2.NewUserService(userRepo)
	eventService := Implementations2.NewEventService(eventRepo)
	contactService := Implementations2.NewContactService(contactRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Router.InitRouter(userService, eventService, contactService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
