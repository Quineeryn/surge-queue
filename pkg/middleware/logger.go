package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		correlationId, _ := r.Context().Value(CorrelationKey).(string)
		if correlationId == "" {
			correlationId = "unknown"
		}
		next.ServeHTTP(w, r)
		log.Printf("x-correlation-id: %s [%s] %s | Duration: %v", correlationId, r.Method, r.URL.Path, time.Since(start))
	})
}
