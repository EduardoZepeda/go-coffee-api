package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/utils"
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
		"/api/v1/login",
		"/api/v1/signup",
	}
)

func IsLoginOrRegisterAttempt(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.HasPrefix(route, p) {
			return true
		}
	}
	return false
}

func AuthenticatedOrReadOnly(app *application.App) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if methodIsSafe(r.Method) || IsLoginOrRegisterAttempt(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			userId, err := utils.GetDataFromToken(r, "userId")
			if err != nil {
				app.Logger.Println(err.Error())
				app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func IsStaffOrReadOnly(app *application.App) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if methodIsSafe(r.Method) || IsLoginOrRegisterAttempt(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			isStaff, err := utils.GetDataFromToken(r, "isStaff")
			if err != nil {
				app.Logger.Println(err)
				app.Respond(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if !isStaff.(bool) {
				app.Respond(w, types.ApiError{Message: "You don't have permission to access this view"}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
func AuthenticatedOnly(app *application.App) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId, err := utils.GetDataFromToken(r, "userId")
			if err != nil {
				app.Logger.Println(err)
				app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
