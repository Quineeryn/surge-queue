package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/module/activity"
	"net/http"
)

func ActivityLogRoute(router *http.ServeMux, svc activity.Service) {
	router.HandleFunc("GET /activity/transfer/{user_id}", handler.GetTransferSummary(svc))
}
