package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/EduardoZepeda/go-coffee-api/models"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"github.com/golang-jwt/jwt/v4"
)

var (
	SAFE_METHODS = []string{
		"GET",
		"OPTIONS",
	}
)

func methodIsSafe(method string) bool {
	for _, safeMethod := range SAFE_METHODS {
		if method == safeMethod {
			return true
		}
	}
	return false
}

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func IsLoginOrRegisterAttempt(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return true
		}
	}
	return false
}

func AuthenticatedOrReadOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methodIsSafe(r.Method) || IsLoginOrRegisterAttempt(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			web.Respond(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
