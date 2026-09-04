package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/module/user"
	"myAPI/pkg/worker"
	"net/http"
)

func UserRoute(router *http.ServeMux, svc user.Service, emailChan chan<- worker.EmailJob) {

	router.HandleFunc("POST /users", handler.CreateUser(svc, emailChan))
	router.HandleFunc("GET /users/{id}", handler.FindByIdUser(svc))
	router.HandleFunc("GET /users", handler.FindAllUser(svc))
	router.HandleFunc("DELETE /users/{id}", handler.DeleteUser(svc))
	router.HandleFunc("PUT /users/{id}", handler.UpdateUser(svc))
	router.HandleFunc("POST /users/transfer", handler.CreateTransfer(svc))

}
