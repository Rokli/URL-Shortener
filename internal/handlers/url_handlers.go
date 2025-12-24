package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func UrlShortener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var long_url string = vars["url"]

	fmt.Println(long_url)
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
}

func GetStats(w http.ResponseWriter, r *http.Request) {

}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {

}
