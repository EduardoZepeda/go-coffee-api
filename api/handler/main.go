package handler

import (
	"log"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/database"
	// Remember to place docs outside of api/handler when deploying in vercel
	// to prevent "Error: Could not find an exported function" error
	_ "github.com/EduardoZepeda/go-coffee-api/docs"
	"github.com/EduardoZepeda/go-coffee-api/handlers"
	"github.com/EduardoZepeda/go-coffee-api/middleware"
	"github.com/EduardoZepeda/go-coffee-api/repository"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Coffee Shops in Gdl API
// @version 1.0
// @description This API returns information about speciality coffee shops in Guadalajara, Mexico.
// @termsOfService http://swagger.io/terms/
// @contact.name Eduardo Zepeda
// @contact.email eduardozepeda@coffeebytes.dev
// @license.name MIT
// @license.url https://mit-license.org/
// @host go-coffee-api.vercel.app
// @BasePath /api/v1
func Api(w http.ResponseWriter, r *http.Request) {
	repo, err := database.NewPostgresRepository()
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	defer repo.Close()
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	api.HandleFunc("/debug_next", handlers.SwaggerDocs).Methods(http.MethodGet)
	api.Use(middleware.RecoverFromPanic)
	api.Use(middleware.RateLimit)
	api.Use(middleware.AuthenticatedOrReadOnly)
	api.HandleFunc("/login", handlers.LoginUser).Methods(http.MethodPost)
	api.HandleFunc("/user", handlers.RegisterUser).Methods(http.MethodPost)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser).Methods(http.MethodGet)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.UpdateUser).Methods(http.MethodPut)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUser).Methods(http.MethodDelete)
	api.HandleFunc("/following/{id:[0-9]+}", handlers.GetUserFollowingAccounts).Methods(http.MethodGet)
	api.HandleFunc("/following", handlers.FollowUser).Methods(http.MethodPost)
	api.HandleFunc("/following/{id:[0-9]+}", handlers.UnfollowUser).Methods(http.MethodDelete)
	api.HandleFunc("/followers/{id:[0-9]+}", handlers.GetUserFollowers).Methods(http.MethodGet)
	api.HandleFunc("/likes", handlers.LikeCoffeeShop).Methods(http.MethodPost)
	api.HandleFunc("/likes", handlers.GetLikedCoffeeShops).Methods(http.MethodGet)
	api.HandleFunc("/likes/{shop_id:[0-9]+}", handlers.UnlikeCoffeeShop).Methods(http.MethodDelete)
	api.Use(middleware.IsStaffOrReadOnly)
	api.HandleFunc("/cafes", handlers.GetCoffeeShops).Methods(http.MethodGet)
	api.HandleFunc("/cafes", handlers.CreateCoffeeShop).Methods(http.MethodPost)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.GetCoffeeShopById).Methods(http.MethodGet)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.UpdateCoffeeShop).Methods(http.MethodPut)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.DeleteCoffeeShop).Methods(http.MethodDelete)
	// the last two cafe endpoints are preserved for compatibility purpose, but its functionality
	// can be replaced by /cafes [get] endpoint
	api.HandleFunc("/cafes/nearest", handlers.GetNearestCoffeeShop).Methods(http.MethodPost)
	api.HandleFunc("/cafes/search/{searchTerm:[a-z]+}", handlers.SearchCoffeeShops).Methods(http.MethodGet)
	api.ServeHTTP(w, r)
}
