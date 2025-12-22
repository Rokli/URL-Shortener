package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This site working.")
	})

	router.HandleFunc("/api/health/db", func(w http.ResponseWriter, r *http.Request) {

	})

	http.Handle("/", router)
	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
