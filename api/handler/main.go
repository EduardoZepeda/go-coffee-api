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
	modifiedHttpSwaggo "github.com/EduardoZepeda/go-coffee-api/modifiedswaggo"
	"github.com/gorilla/mux"
)

var app *application.App
var err error

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
	app, err = application.New()
	if err != nil {
		log.Fatal("Server couldn't start")
	}
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.RecoverFromPanic(app), middleware.CorsAllowAll(app), middleware.RateLimit(app))
	// api.PathPrefix("/ws").Handler(handlers.HandleWebSockets(app))
	api.PathPrefix("/swagger").Handler(modifiedHttpSwaggo.WrapHandler)
	api.PathPrefix("/healthcheck").Handler(handlers.Healtcheck(app)).Methods(http.MethodGet)
	loginRegisterApi := api.PathPrefix("/").Subrouter()
	loginRegisterApi.Use(middleware.AuthenticatedOrReadOnly(app))
	loginRegisterApi.HandleFunc("/login", handlers.LoginUser(app)).Methods(http.MethodPost)
	loginRegisterApi.HandleFunc("/signup", handlers.RegisterUser(app)).Methods(http.MethodPost)
	loginRegisterApi.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser(app)).Methods(http.MethodGet)
	loginRegisterApi.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser(app)).Methods(http.MethodPut)
	loginRegisterApi.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser(app)).Methods(http.MethodDelete)
	followersAndLikes := api.PathPrefix("/").Subrouter()
	// Likes and following are only available to authenticated users
	followersAndLikes.Use(middleware.AuthenticatedOnly(app))
	followersAndLikes.HandleFunc("/following/{id:[0-9]+}", handlers.GetUserFollowingAccounts(app)).Methods(http.MethodGet)
	followersAndLikes.HandleFunc("/following", handlers.FollowUser(app)).Methods(http.MethodPost)
	followersAndLikes.HandleFunc("/following/{id:[0-9]+}", handlers.UnfollowUser(app)).Methods(http.MethodDelete)
	followersAndLikes.HandleFunc("/followers/{id:[0-9]+}", handlers.GetUserFollowers(app)).Methods(http.MethodGet)
	followersAndLikes.HandleFunc("/likes", handlers.LikeCoffeeShop(app)).Methods(http.MethodPost)
	followersAndLikes.HandleFunc("/likes", handlers.GetLikedCoffeeShops(app)).Methods(http.MethodGet)
	followersAndLikes.HandleFunc("/likes/{shop_id:[0-9]+}", handlers.UnlikeCoffeeShop(app)).Methods(http.MethodDelete)
	// Coffee shops endpoints, this routes are protected, and only staff members can use unsafe methods
	coffeeShopsApi := api.PathPrefix("/coffee-shops").Subrouter()
	coffeeShopsApi.Use(middleware.IsStaffOrReadOnly(app))
	coffeeShopsApi.HandleFunc("", handlers.GetCoffeeShops(app)).Methods(http.MethodGet)
	coffeeShopsApi.HandleFunc("", handlers.CreateCoffeeShop(app)).Methods(http.MethodPost)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}", handlers.GetCoffeeShopById(app)).Methods(http.MethodGet)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}", handlers.UpdateCoffeeShop(app)).Methods(http.MethodPut)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}", handlers.DeleteCoffeeShop(app)).Methods(http.MethodDelete)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}/coffee-bags", handlers.GetCoffeeBagByCoffeeShop(app)).Methods(http.MethodGet)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}/coffee-bags/{coffee_bag_id:[0-9]+}", handlers.AddCoffeeBagToCoffeeShop(app)).Methods(http.MethodPost)
	coffeeShopsApi.HandleFunc("/{id:[0-9]+}/coffee-bags/{coffee_bag_id:[0-9]+}", handlers.DeleteCoffeeBagFromCoffeeShop(app)).Methods(http.MethodDelete)

	// Coffee bags endpoints, this routes are protected, and only staff members can use unsafe methods
	coffeeBagsApi := api.PathPrefix("/coffee-bags").Subrouter()
	coffeeBagsApi.Use(middleware.IsStaffOrReadOnly(app))
	coffeeBagsApi.HandleFunc("", handlers.CreateCoffeeBag(app)).Methods(http.MethodPost)
	coffeeBagsApi.HandleFunc("", handlers.GetCoffeeBags(app)).Methods(http.MethodGet)
	coffeeBagsApi.HandleFunc("/{id:[0-9]+}", handlers.GetCoffeeBagById(app)).Methods(http.MethodGet)
	coffeeBagsApi.HandleFunc("/{id:[0-9]+}", handlers.UpdateCoffeeBag(app)).Methods(http.MethodPut)
	coffeeBagsApi.HandleFunc("/{id:[0-9]+}", handlers.DeleteCoffeeBag(app)).Methods(http.MethodDelete)

	// Feed for user, only authenticated users can access it
	feedApi := api.PathPrefix("/feed").Subrouter()
	feedApi.Use(middleware.AuthenticatedOnly(app))
	feedApi.HandleFunc("", handlers.GetUserFeed(app)).Methods(http.MethodGet)
	router.ServeHTTP(w, r)
}
