package handler

import (
	"myAPI/pkg/adapter/presenter"
	"myAPI/pkg/entity/query"
	"myAPI/pkg/middleware"
	"myAPI/pkg/module/activity"
	"myAPI/pkg/shared/response"
	"myAPI/pkg/shared/validator"
	"net/http"

	"github.com/google/uuid"
)

func GetTransferSummary(svc activity.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(query.TransferSummary)

		userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
		if !ok {
			response.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.UserID = userID.String()
		if err := validator.Validate(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		res, err := svc.GetTransferSummary(r.Context(), req)
		if err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := presenter.NewTransferSummaryOutput(res)
		response.Success(w, http.StatusOK, "Transfer Summary", resp)
	}
}
