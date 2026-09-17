package route

import (
	"myAPI/pkg/adapter/handler"
	"myAPI/pkg/middleware"
	"myAPI/pkg/module/activity"
	"myAPI/pkg/security"
	"net/http"
)

func ActivityLogRoute(router *http.ServeMux, svc activity.Service, jwtManager *security.JWTManager) {
	router.HandleFunc("GET /transfers/summary", middleware.Authenticate(jwtManager, handler.GetTransferSummary(svc)))
}
