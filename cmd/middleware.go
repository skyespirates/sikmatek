package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/skyespirates/sikmatek/internal/utils"
)

func (app *application) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.LogInfo(r, "request")
		next.ServeHTTP(w, r)
		app.logger.LogInfo(r, "response")
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
