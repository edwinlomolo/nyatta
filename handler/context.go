package handler

import (
	"context"
	"net/http"
)

// AddContext - feed custom context onto any request handler through context
func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
