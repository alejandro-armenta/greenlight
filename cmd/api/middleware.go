package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			w.Header().Add("Vary", "Authorization")

			authorizationHeader := r.Header.Get("Authorization")

			if authorizationHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			headerParts := strings.Split(authorizationHeader, " ")

			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				app.invalidAuthenticationTokenResponse()
				return
			}

			ale := headerParts[1]

			next.ServeHTTP(w, r)
		})
}

func (app *application) rateLimit(next http.Handler) http.Handler {

	if !app.config.limiter.enabled {
		return next
	}

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      = sync.Mutex{}
		clients = map[string]*client{}
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			ip := realip.FromRequest(r)

			mu.Lock()

			if _, found := clients[ip]; !found {
				clients[ip] = &client{
					limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst),
				}
			}

			clients[ip].lastSeen = time.Now()

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceedResponse(w, r)
				return
			}

			mu.Unlock()

			next.ServeHTTP(w, r)
		})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				pv := recover()

				if pv != nil {
					//esto va a cerrar la madre
					w.Header().Set("Connection", "close")
					app.serverErrorResponse(w, r, fmt.Errorf("%v", pv))
				}
			}()

			next.ServeHTTP(w, r)
		})
}
