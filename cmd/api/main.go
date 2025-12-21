package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
		fmt.Fprintln(w, "This site working.")
	})

	http.Handle("/", router)
	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
