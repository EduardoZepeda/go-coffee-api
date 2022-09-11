package handlers

import (
	"net/http"
	"regexp"

	"github.com/EduardoZepeda/go-coffee-api/web"
)

type DebugNext struct {
	Uri     string
	Matches []string
}

func SwaggerDocs(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^(.*/)([^?].*)?[?|.]*$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	web.Respond(w, DebugNext{Uri: r.URL.Path, Matches: matches}, http.StatusOK)
}
