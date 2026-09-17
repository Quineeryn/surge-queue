package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/middleware"
	"myAPI/pkg/module/user"
	"myAPI/pkg/security"
	"myAPI/pkg/worker"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func UserRoute(router *http.ServeMux, svc user.Service, emailChan chan<- worker.EmailJob, jwtManager *security.JWTManager, redis *redis.Client) {

	router.HandleFunc("POST /users", middleware.RateLimiter(redis, 5, 5*time.Minute)(handler.CreateUser(svc, emailChan)))
	router.HandleFunc("GET /users/{id}", handler.FindByIdUser(svc))

	protected := func(pattern string, h http.HandlerFunc) {
		router.HandleFunc(pattern, middleware.Authenticate(jwtManager, h))
	}

	protected("GET /users", handler.FindAllUser(svc))
	protected("DELETE /users/{id}", handler.DeleteUser(svc))
	protected("PUT /users/{id}", handler.UpdateUser(svc))
	protected("POST /transfers", middleware.Idempotency(redis, handler.CreateTransfer(svc)))
}
