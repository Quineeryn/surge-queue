package handler

import (
	"context"
	"encoding/json"
	"errors"
	"myAPI/pkg/adapter/presenter"
	"myAPI/pkg/entity"
	"myAPI/pkg/module/ping"
	"net/http"
)

func GetPongMessage(svc ping.PingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := svc.GetMessage(r.Context())

		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Gateway Timout", http.StatusGatewayTimeout)
			return
		}

		outputs := make([]presenter.PingOutput, 0, len(res))

		for _, item := range res {
			outputs = append(outputs, *presenter.NewPingOutput(&item))
		}
		json.NewEncoder(w).Encode(outputs)
	}
}

func CreateMessage(svc ping.PingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(entity.PingInput)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		dto := entity.NewPingDtoFromInput(req)

		res := svc.CreateMessage(dto)

		resp := presenter.NewPingOutput(res)

		json.NewEncoder(w).Encode(resp)
	}
}
