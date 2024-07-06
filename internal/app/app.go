package app

import (
	"eventPlanner/internal/repository"
	"eventPlanner/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"eventPlanner/internal/database"

	"context"
    "encoding/json"
    
    "eventPlanner/internal/models"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "net/http"
)

func RunApp(port string) {
	//cfg := config.Config{}
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

    logrus.Info("Starting server on :8082")
    

	userService := service.NewUserService(userRepo)

	eventRepo := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepo)

	contactRepo := repository.NewContactRepository()
	contactService := service.NewContactService(contactRepo)
	// Ручки
	e.POST("/auth/register", userService.Register)
	e.POST("/auth/login", userService.Login)

	e.POST("/events/create/:id", eventService.CreateEvent)
	e.GET("/events/:id", eventService.GetAllEvents)

	e.POST("/contacts", contactService.CreateContact)
	e.GET("/contacts", contactService.GetAllContacts)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
