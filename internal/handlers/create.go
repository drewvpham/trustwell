package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "my-rest-service/internal/models"
    "my-rest-service/internal/store"
    "github.com/google/uuid"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
    var event models.Event
    err := json.NewDecoder(r.Body).Decode(&event)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Validate required fields
    if event.Type != "shipping" && event.Type != "receiving" {
        http.Error(w, "Invalid event type", http.StatusBadRequest)
        return
    }

    // Derive fields
    event.ID = uuid.New().String()
    event.CreatedAt = time.Now().UTC()
    event.CreatedBy = r.Context().Value("userID").(string)
    event.IsDeleted = false

    // Store event
    store.AddEvent(event)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(event)
}
