package middleware

import (
	"fmt"
	"myAPI/pkg/shared/response"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client, limit int64, window time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			host := getClientIP(r)
			key := fmt.Sprintf("rate_limit:%s", host)
			now := time.Now().UnixNano()
			windowStart := now - window.Nanoseconds()

			res, err := rateLimitScript.Run(r.Context(), rdb, []string{key}, now, windowStart, limit, int(window.Seconds())).Int()
			if err != nil || res == 0 {
				if res == 0 {
					response.Error(w, "Too many requests", http.StatusTooManyRequests)
					return
				}
				response.Error(w, "Rate limit exceeded. Please try again later.", http.StatusTooManyRequests)
				next(w, r)
			}
			next(w, r)

		}
	}
}

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ip := strings.Split(xff, ",")[0]
		return strings.TrimSpace(ip)
	}

	xri := r.Header.Get("X-Real-IP")

	if xri != "" {
		return xri
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	return host

}

var rateLimitScript = redis.NewScript(`
	redis.call('ZREMRANGEBYSCORE', KEYS[1], '-inf', ARGV[2])
	local current = redis.call('ZCARD', KEYS[1])
	if current < tonumber(ARGV[3]) then
		redis.call('ZADD', KEYS[1], ARGV[1], ARGV[1])
		redis.call('EXPIRE', KEYS[1], ARGV[4])
		return 1
	end
	return 0
`)
