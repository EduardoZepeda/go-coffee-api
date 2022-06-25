package handler

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/middleware"
	"github.com/EduardoZepeda/go-coffee-api/views"
	"github.com/gorilla/mux"
)

func Api(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.AuthenticatedOrReadOnly)
	api.HandleFunc("/cafes", views.GetCafes).Methods(http.MethodGet)
	api.HandleFunc("/cafes", views.CreateCafe).Methods(http.MethodPost)
	api.HandleFunc("/cafes/{id}", views.GetCafeById).Methods(http.MethodGet)
	api.HandleFunc("/cafes/{id}", views.UpdateCafe).Methods(http.MethodPut)
	api.HandleFunc("/cafes/{id}", views.DeleteCafe).Methods(http.MethodDelete)
	api.HandleFunc("/search/{searchTerm}", views.SearchCafe).Methods(http.MethodGet)
	api.HandleFunc("/login", views.LoginUser).Methods(http.MethodPost)
	api.HandleFunc("/signup", views.SignupUser).Methods(http.MethodPost)
	api.ServeHTTP(w, r)
}
