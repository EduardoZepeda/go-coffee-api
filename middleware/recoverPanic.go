package middleware

import (
	"net/http"

	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/web"
)

func RecoverFromPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This anonymous function will always run
		defer func() {
			// Use the builtin recover function
			if err := recover(); err != nil {
				// close the current connection
				w.Header().Set("Connection", "close")
				// Send an error response to the user
				web.Respond(w, types.ApiError{Message: "Closed connection"}, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
