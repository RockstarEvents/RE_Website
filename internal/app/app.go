package app

import (
	"context"
	"encoding/json"
	"eventPlanner/internal/database"
	"eventPlanner/internal/models"
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunApp(port string) {
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

	router := mux.NewRouter()
	router.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userID, err := userRepo.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("User created with ID: %d", userID)))
	}).Methods("POST")


	userService := service.NewUserService(userRepo)
	eventService := service.NewEventService(eventRepo)
	contactService := service.NewContactService(contactRepo) // Инициализируем сервис контактов с новым репозиторием

	// Ручки
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create", eventService.CreateEvent)
	e.GET("/events", eventService.GetAllEvents)

	e.GET("/contacts", contactService.GetAllContacts)       // Добавляем ручку для получения всех контактов
	e.POST("/contacts/select", contactService.SelectContacts) // Добавляем ручку для выбора контактов

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
