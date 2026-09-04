package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey string

const CorrelationKey ctxKey = "x-correlation-id"

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("x-correlation-id")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}
		w.Header().Set("x-correlation-id", correlationID)
		ctx := context.WithValue(r.Context(), CorrelationKey, correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
