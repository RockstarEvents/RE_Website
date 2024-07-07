package repository

import (
    "context"
    "errors"
    "eventPlanner/internal/models"
    "github.com/jackc/pgx/v4"
    "github.com/sirupsen/logrus"
)

type EventRepository interface {
    CreateEvent(event models.Event) error
    GetAllEvents() ([]models.Event, error)
}

type PostgresEventRepository struct {
    conn *pgx.Conn
}

func NewEventRepository(conn *pgx.Conn) EventRepository {
    return &PostgresEventRepository{conn: conn}
}

func (r *PostgresEventRepository) CreateEvent(event models.Event) error {
    query := `INSERT INTO events (name, shape, place, begin_time, duration) VALUES ($1, $2, $3, $4, $5)`
    _, err := r.conn.Exec(context.Background(), query, event.NameEvent, event.Shape, event.Place, event.BeginTime, event.Duration)
    if err != nil {
        logrus.Errorf("Error creating event: %v", err)
        return err
    }
    return nil
}

func (r *PostgresEventRepository) GetAllEvents() ([]models.Event, error) {
    query := `SELECT id, name, shape, place, begin_time, duration FROM events`
    rows, err := r.conn.Query(context.Background(), query)
    if err != nil {
        logrus.Errorf("Error fetching events: %v", err)
        return nil, err
    }
    defer rows.Close()

    var events []models.Event
    for rows.Next() {
        var event models.Event
        err := rows.Scan(&event.ID, &event.NameEvent, &event.Shape, &event.Place, &event.BeginTime, &event.Duration)
        if err != nil {
            logrus.Errorf("Error scanning event: %v", err)
            return nil, err
        }
        events = append(events, event)
    }

    if rows.Err() != nil {
        logrus.Errorf("Row error: %v", rows.Err())
        return nil, rows.Err()
    }

    if len(events) == 0 {
        return nil, errors.New("no events found")
    }

    return events, nil
}
