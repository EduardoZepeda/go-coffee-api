package handlers

import (
	"net/http"
	"os"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/models"
)

// Healthcheck endping
// @Summary      Returns the server status
// @Description  Returns the api version, the environment and the server status
// @Tags         healthcheck
// @Success      200  {object}  models.HealtcheckResponse
// @Router       /healthcheck [get]
func Healtcheck(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &models.HealtcheckResponse{
			Version:     "1.0",
			Status:      "up",
			Environment: os.Getenv("DB_USER"),
		}
		app.Respond(w, response, http.StatusOK)
	}
}
