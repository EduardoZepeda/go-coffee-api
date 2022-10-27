package handler

import (
	"log"
	"net/http"

	// Remember to place docs outside of api/handler when deploying in vercel
	// to prevent "Error: Could not find an exported function" error

	"github.com/EduardoZepeda/go-coffee-api/application"
	_ "github.com/EduardoZepeda/go-coffee-api/docs"
	"github.com/EduardoZepeda/go-coffee-api/handlers"
	"github.com/EduardoZepeda/go-coffee-api/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var app *application.App
var err error

func init() {
	app, err = application.New()
	if err != nil {
		log.Fatal("Server couldn't start")
	}
}

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
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.RecoverFromPanic(app))
	api.Use(middleware.RateLimit(app))
	api.Use(middleware.AuthenticatedOrReadOnly(app))
	api.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	api.HandleFunc("/login", handlers.LoginUser(app)).Methods(http.MethodPost)
	api.HandleFunc("/user", handlers.RegisterUser(app)).Methods(http.MethodPost)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.GetUser(app)).Methods(http.MethodGet)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.UpdateUser(app)).Methods(http.MethodPut)
	api.HandleFunc("/user/{id:[0-9]+}", handlers.DeleteUser(app)).Methods(http.MethodDelete)
	api.HandleFunc("/following/{id:[0-9]+}", handlers.GetUserFollowingAccounts(app)).Methods(http.MethodGet)
	api.HandleFunc("/following", handlers.FollowUser(app)).Methods(http.MethodPost)
	api.HandleFunc("/following/{id:[0-9]+}", handlers.UnfollowUser(app)).Methods(http.MethodDelete)
	api.HandleFunc("/followers/{id:[0-9]+}", handlers.GetUserFollowers(app)).Methods(http.MethodGet)
	api.HandleFunc("/likes", handlers.LikeCoffeeShop(app)).Methods(http.MethodPost)
	api.HandleFunc("/likes", handlers.GetLikedCoffeeShops(app)).Methods(http.MethodGet)
	api.HandleFunc("/likes/{shop_id:[0-9]+}", handlers.UnlikeCoffeeShop(app)).Methods(http.MethodDelete)
	api.Use(middleware.IsStaffOrReadOnly(app))
	api.HandleFunc("/cafes", handlers.GetCoffeeShops(app)).Methods(http.MethodGet)
	api.HandleFunc("/cafes", handlers.CreateCoffeeShop(app)).Methods(http.MethodPost)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.GetCoffeeShopById(app)).Methods(http.MethodGet)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.UpdateCoffeeShop(app)).Methods(http.MethodPut)
	api.HandleFunc("/cafes/{id:[0-9]+}", handlers.DeleteCoffeeShop(app)).Methods(http.MethodDelete)
	// the last two cafe endpoints are preserved for compatibility purpose, but its functionality
	// can be replaced by /cafes [get] endpoint
	api.HandleFunc("/cafes/nearest", handlers.GetNearestCoffeeShop(app)).Methods(http.MethodPost)
	api.HandleFunc("/cafes/search/{searchTerm:[a-z]+}", handlers.SearchCoffeeShops(app)).Methods(http.MethodGet)
	api.ServeHTTP(w, r)
}
