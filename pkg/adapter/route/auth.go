package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/module/auth"
	"net/http"
)

func AuthRoute(router *http.ServeMux, svc auth.Service) {
	router.HandleFunc("POST /login", handler.Login(svc))
}
