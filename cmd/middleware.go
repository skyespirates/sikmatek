package main

import (
	"fmt"
	"net"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/utils"
	"golang.org/x/time/rate"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (app *application) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rec := &responseRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		app.logger.PrintInfo("request_started", map[string]string{
			"method":      r.Method,
			"path":        r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		})

		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		durationMs := float64(duration) / float64(time.Millisecond)

		app.logger.PrintInfo("request_completed", map[string]string{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status":      strconv.Itoa(rec.status),
			"duration":    fmt.Sprintf("%.2fms", durationMs),
			"remote_addr": r.RemoteAddr,
		})
	})
}

func (app *application) authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ambil header authorization
		authorizationToken := r.Header.Get("Authorization")
		if authorizationToken == "" {
			response.Error(w, http.StatusUnauthorized, "missing token", nil)
			return
		}

		// split it
		parts := strings.Split(authorizationToken, " ")
		if parts[0] != "Bearer" {
			response.Error(w, http.StatusUnauthorized, "token must be bearer", nil)
			return
		}

		// ambil yg tokennya aja
		token := parts[1]

		// verify token
		claim, err := utils.VerifyToken(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid credentials", err)
			return
		}

		r = utils.ContextSetUser(r, claim)

		next.ServeHTTP(w, r)
	})
}

func (app *application) authorize(allowedRoles ...int) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			claim := utils.ContextGetUser(r.Context()) // retrieve claims from context
			if claim == nil {
				response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
				return
			}

			if slices.Contains(allowedRoles, claim.RoleId) {
				next.ServeHTTP(w, r)
				return
			}

			response.Error(w, http.StatusForbidden, "forbidden", nil)
		}
	}

}

func (app *application) rateLimit(next http.Handler) http.Handler {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// clean up expired client
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			app.logger.PrintError(err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		mu.Lock()

		_, ok := clients[ip]
		if !ok {
			clients[ip] = &client{
				limiter: rate.NewLimiter(2, 10),
			}
		}

		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			app.logger.PrintInfo("request_rate_limited", map[string]string{
				"method": r.Method,
				"path":   r.URL.Path,
				"ip":     ip,
			})
			response.Error(w, http.StatusTooManyRequests, "rate limit exceeded", nil)
			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
