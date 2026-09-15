package handler

import (
	"encoding/json"
	"myAPI/pkg/adapter/presenter"
	"myAPI/pkg/entity"
	"myAPI/pkg/module/auth"
	"myAPI/pkg/shared/response"
	"myAPI/pkg/shared/validator"
	"net/http"
)

func Login(svc auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(entity.LoginInput)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		res, err := svc.Login(r.Context(), req)

		if err != nil {
			response.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		resp := presenter.NewLoginOutput(res)

		response.Success(w, http.StatusOK, "Login successfull", resp)
	}
}
