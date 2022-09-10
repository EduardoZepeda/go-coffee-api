package handlers

import (
	"net/http"
	"regexp"

	"github.com/EduardoZepeda/go-coffee-api/web"
)

func SwaggerDocs(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^(.*/)([^?].*)?[?|.]*$`)
	matches := re.FindStringSubmatch(r.RequestURI)
	web.Respond(w, matches, http.StatusOK)
}
