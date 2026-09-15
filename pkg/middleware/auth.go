package middleware

import (
	"context"
	"myAPI/pkg/security"
	"myAPI/pkg/shared/response"
	"net/http"
	"strings"
)

type authKey string

const (
	UserIDKey authKey = "user_id"
	RoleKey   authKey = "role"
)

func Authenticate(jwtManager *security.JWTManager, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(w, "Unauthorized", http.StatusUnauthorized)
			return

		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			response.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, RoleKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
