package utils

import (
	"context"
	"log"
	"net/http"
)

type contextKey string

var UserContextKey = contextKey("user")

func ContextSetUser(r *http.Request, claim *Claims) *http.Request {
	ctx := context.WithValue(r.Context(), UserContextKey, claim)
	return r.WithContext(ctx)
}

func ContextGetUser(ctx context.Context) *Claims {
	claim, ok := ctx.Value(UserContextKey).(*Claims)
	if !ok {
		log.Fatal("missing user value in request context")
	}

	return claim
}
