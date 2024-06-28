package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "my-rest-service/internal/store"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    userID := r.Context().Value("userID").(string)
    err := store.SoftDeleteEvent(id, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
