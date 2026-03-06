package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

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
			res := map[string]string{}
			res["status"] = "unauthorized"
			res["message"] = "missing token"
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		}

		// split it
		parts := strings.Split(authorizationToken, " ")
		if parts[0] != "Bearer" {
			res := map[string]string{}
			res["status"] = "unauthorized"
			res["message"] = "token must be Bearer"
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		}

		// ambil yg tokennya aja
		token := parts[1]

		// verify token
		claim, err := utils.VerifyToken(token)
		if err != nil {
			res := map[string]string{}
			res["status"] = "invalid credential"
			res["message"] = err.Error()
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		}

		r = utils.ContextSetUser(r, claim)

		next.ServeHTTP(w, r)
	})
}

func (app *application) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) authorize(allowedRoles ...int) func(http.HandlerFunc) http.HandlerFunc {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			claim := utils.ContextGetUser(r.Context()) // retrieve claims from context
			if claim == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			if slices.Contains(allowedRoles, claim.RoleId) {
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "forbidden", http.StatusForbidden)
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
				limiter: rate.NewLimiter(2, 4),
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
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("rate limit exceeded"))
			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
