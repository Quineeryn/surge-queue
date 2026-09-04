package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/module/ping"
	"net/http"
)

func PingRoute(mux *http.ServeMux, svc ping.PingService) {
	mux.HandleFunc("/ping", handler.GetPongMessage(svc))
	mux.HandleFunc("/pong", handler.CreateMessage(svc))
}
