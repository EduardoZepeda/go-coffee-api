package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/database"
	"github.com/gorilla/mux"
)

type App struct {
	Repo   *database.PostgresRepository
	Router *mux.Router
	Logger *log.Logger
}

func (app *App) Respond(w http.ResponseWriter, data interface{}, statusCode int) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}
	return nil
}

func (app *App) SetPostgresRepository() error {
	repo, err := database.NewPostgresRepository()
	if err != nil {
		log.Fatal(err)
	}
	app.Logger.Println("Initialized database")
	app.Repo = repo
	return nil
}

func (app *App) SetRouter(router *mux.Router) error {
	app.Router = router
	return nil
}

func (app *App) SetLogger() error {
	// Default logger for now
	app.Logger = log.Default()
	app.Logger.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	app.Logger.Println("Logging events in application")
	return nil
}

func (app *App) Initialize() error {
	err := app.SetLogger()
	if err != nil {
		return err
	}
	err = app.SetPostgresRepository()
	if err != nil {
		app.Logger.Fatal(err)
		return err
	}
	app.Logger.Println("App Initialized")
	return nil
}

func New() (*App, error) {
	newApp := App{}
	err := newApp.Initialize()
	if err != nil {
		return nil, err
	}
	return &newApp, nil
}
