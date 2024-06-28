package store

import (
    "errors"
    "sync"
    "my-rest-service/internal/models"
)

var (
    mu     sync.Mutex
    events = make(map[string]models.Event)
)

func AddEvent(event models.Event) {
    mu.Lock()
    defer mu.Unlock()
    events[event.ID] = event
}

func SoftDeleteEvent(id, userID string) error {
    mu.Lock()
    defer mu.Unlock()
    event, exists := events[id]
    if !exists || event.IsDeleted || event.CreatedBy != userID {
        return errors.New("Event not found")
    }
    event.IsDeleted = true
    events[id] = event
    return nil
}

func GetEvent(id, userID string) (models.Event, error) {
    mu.Lock()
    defer mu.Unlock()
    event, exists := events[id]
    if !exists || event.IsDeleted || event.CreatedBy != userID {
        return event, errors.New("Event not found")
    }
    return event, nil
}

func ListEvents(userID string) []models.Event {
    mu.Lock()
    defer mu.Unlock()
    userEvents := []models.Event{}
    for _, event := range events {
        if event.CreatedBy == userID && !event.IsDeleted {
            userEvents = append(userEvents, event)
        }
    }
    return userEvents
}
