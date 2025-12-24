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

	router.HandleFunc("/api/v1/get/short/{url}", handlers.UrlShortener)

	http.Handle("/", router)
	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
