package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "my-rest-service/internal/handlers"
    "my-rest-service/internal/middleware"
)

func main() {
    router := mux.NewRouter()
    router.Use(middleware.AuthMiddleware)

    router.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
    router.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")
    router.HandleFunc("/events/{id}", handlers.GetEvent).Methods("GET")
    router.HandleFunc("/events", handlers.ListEvents).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}
