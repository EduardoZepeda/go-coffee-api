package handler

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/api/cafes"
	"github.com/gorilla/mux"
)

func Api(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/api/cafes", cafes.GetCafes).Methods(http.MethodGet)
	router.HandleFunc("/api/cafes/{id}", cafes.GetCafeById).Methods(http.MethodGet)
	router.HandleFunc("/api/cafes/search/{searchTerm}", cafes.SearchCafe).Methods(http.MethodGet)
	router.ServeHTTP(w, r)
}
