package middleware

import (
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/application"
	"github.com/EduardoZepeda/go-coffee-api/types"
	"golang.org/x/time/rate"
)

func RateLimit(app *application.App) func(h http.Handler) http.Handler {
	type client struct {
		limiter     *rate.Limiter
		lastRequest time.Time
	}

	var (
		mu sync.Mutex
		// Create a black list for clients with too much requests
		clients = make(map[string]*client)
	)
	// go routine that deletes the go through the black list every 5 minutes and deletes all the inactive clients
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			mu.Lock()
			// If the client last request was more than 2 minute ago, remove them from the black list
			for ip, client := range clients {
				if time.Since(client.lastRequest) > 2*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ip string
			var err error
			if os.Getenv("MODE") == "dev" {
				// The following code works allow to handle ip addresses that include the port, ideal for localhost
				ip, _, err = net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					app.Logger.Println(err)
					app.Respond(w, types.ApiError{Message: err.Error()}, http.StatusInternalServerError)
				}
			} else {
				// ParseIP doesn't allow the port to be included in the url, therefore, if we're in production we use ParseIP instead
				parsedIp := net.ParseIP(r.RemoteAddr)
				if parsedIp == nil {
					app.Respond(w, types.ApiError{Message: "Couldn't parse your ip address"}, http.StatusInternalServerError)
					return
				}
				ip = parsedIp.String()
			}
			mu.Lock()
			if _, found := clients[ip]; !found {
				// If the ip doesn't exist, add it to the black list
				clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
			}
			// Update the last request for the client.
			clients[ip].lastRequest = time.Now()
			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.Respond(w, types.ApiError{Message: "You're making too much requests. Please wait some time before trying again."}, http.StatusTooManyRequests)
				return
			}
			mu.Unlock()
			next.ServeHTTP(w, r)
		})
	}
}
