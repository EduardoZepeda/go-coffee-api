package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/EduardoZepeda/go-coffee-api/types"
	"github.com/EduardoZepeda/go-coffee-api/web"
	"golang.org/x/time/rate"
)

func RateLimit(next http.Handler) http.Handler {
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedIp := net.ParseIP(r.RemoteAddr)
		if parsedIp == nil {
			web.Respond(w, types.ApiError{Message: "Couldn't parse your ip address"}, http.StatusInternalServerError)
			return
		}
		ip := parsedIp.String()
		mu.Lock()

		if _, found := clients[ip]; !found {
			// If the ip doesn't exist, add it to the black list
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		// Update the last request for the client.
		clients[ip].lastRequest = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			web.Respond(w, types.ApiError{Message: "You're making too much requests. Please wait some time before trying again."}, http.StatusTooManyRequests)
			return
		}
		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
