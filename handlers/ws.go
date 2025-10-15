package handlers

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// logica interna
		return true
	},
}

func HandleWebSockets(app *application.App) http.HandlerFunc {
	app.Logger.Println("Connected")
	return app.Hub.HandleWebSocket
}
