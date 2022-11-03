package middleware

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/application"
)

func CorsAllowAll(app *application.App) func(http http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}
}
