package handlers

import (
	"net/http"
	"regexp"

	"github.com/EduardoZepeda/go-coffee-api/application"
)

type DebugNext struct {
	Uri     string
	Matches []string
}

func SwaggerDocs(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile(`^(.*/)([^?].*)?[?|.]*$`)
		matches := re.FindStringSubmatch(r.RequestURI)
		app.Respond(w, DebugNext{Uri: r.RequestURI, Matches: matches}, http.StatusOK)
	}
}
