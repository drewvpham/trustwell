package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "my-rest-service/internal/store"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    userID := r.Context().Value("userID").(string)
    event, err := store.GetEvent(id, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(event)
}
