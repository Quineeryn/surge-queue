package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func Idempotency(redis *redis.Client, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-Idempotency-Key")
		userID, _ := r.Context().Value(UserIDKey).(string)

		if key == "" || userID == "" {
			next(w, r)
			return
		}

		redisKey := fmt.Sprintf("idempotency:%s:%s", userID, key)
		cachedData, err := redis.Get(r.Context(), redisKey).Result()

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedData))
			return
		}
		rec := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           new(bytes.Buffer),
		}
		next(rec, r)
		if rec.statusCode >= 200 && rec.statusCode < 300 {
			redis.Set(r.Context(), redisKey, rec.body.Bytes(), 24*time.Hour)
		}
	}
}

func (s *responseRecorder) WriteHeader(code int) {
	s.statusCode = code
	s.ResponseWriter.WriteHeader(code)
}

func (s *responseRecorder) Write(b []byte) (int, error) {
	if _, err := s.body.Write(b); err != nil {
		return 0, err
	}

	if s.statusCode == 0 {
		s.statusCode = http.StatusOK
	}
	return s.ResponseWriter.Write(b)
}
