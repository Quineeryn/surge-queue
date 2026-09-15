package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/middleware"
	"myAPI/pkg/module/user"
	"myAPI/pkg/security"
	"myAPI/pkg/worker"
	"net/http"
)

func UserRoute(router *http.ServeMux, svc user.Service, emailChan chan<- worker.EmailJob, jwtManager *security.JWTManager) {

	router.HandleFunc("POST /users", handler.CreateUser(svc, emailChan))
	router.HandleFunc("GET /users/{id}", handler.FindByIdUser(svc))

	protected := func(pattern string, h http.HandlerFunc) {
		router.HandleFunc(pattern, middleware.Authenticate(jwtManager, h))
	}

	protected("GET /users", handler.FindAllUser(svc))
	protected("DELETE /users/{id}", handler.DeleteUser(svc))
	protected("PUT /users/{id}", handler.UpdateUser(svc))
	protected("POST /users/transfer", handler.CreateTransfer(svc))
}
