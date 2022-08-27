package handler

import (
	"log"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/database"
	"github.com/EduardoZepeda/go-coffee-api/handlers"
	"github.com/EduardoZepeda/go-coffee-api/middleware"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/gorilla/mux"
)

func init() {
	repo, err := database.NewPostgresRepository()
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}

func Api(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.AuthenticatedOrReadOnly)
	api.HandleFunc("/cafes", handlers.GetCafes).Methods(http.MethodGet)
	api.HandleFunc("/cafes", handlers.CreateCafe).Methods(http.MethodPost)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.GetCafeById).Methods(http.MethodGet)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.UpdateCafe).Methods(http.MethodPut)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.DeleteCafe).Methods(http.MethodDelete)
	// We prefer a post request to prevent user's location getting saved as links on databases
	api.HandleFunc("/cafes/nearest", handlers.GetNearestCafes).Methods(http.MethodPost)
	api.HandleFunc("/cafes/search/{searchTerm:[a-z]+}", handlers.SearchCafe).Methods(http.MethodGet)
	api.HandleFunc("/login", handlers.LoginUser).Methods(http.MethodPost)
	api.HandleFunc("/signup", handlers.SignupUser).Methods(http.MethodPost)
	api.ServeHTTP(w, r)
}
