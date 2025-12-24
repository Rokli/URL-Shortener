package main

import (
	"fmt"
	"net/http"

	"github.com/Rokli/URL-Shortener/internal/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/health", handlers.HealthCheckAll).Methods("GET")

	router.HandleFunc("/{shortCode}", handlers.RedirectUrl).Methods("GET")

	router.HandleFunc("/api/v1/urls/{code}/stats", handlers.GetStats).Methods("GET")

	router.HandleFunc("/api/v1/urls/{code}", handlers.DeleteUrl).Methods("DELETE")

	router.HandleFunc("/api/v1/get/shorten/{url}", handlers.UrlShortener).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
