package handlers

import (
    "encoding/json"
    "net/http"
    "my-rest-service/internal/store"
)

func ListEvents(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(string)
    events := store.ListEvents(userID)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(events)
}
