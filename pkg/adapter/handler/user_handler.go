package handler

import (
	"encoding/json"
	"errors"
	"log"
	"myAPI/pkg/adapter/presenter"
	"myAPI/pkg/entity"
	"myAPI/pkg/module/user"
	"myAPI/pkg/shared/response"
	"myAPI/pkg/shared/validator"
	"myAPI/pkg/worker"
	"net/http"
)

func CreateUser(svc user.Service, emailChan chan<- worker.EmailJob) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(entity.UserInput)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		dto := entity.NewUserDtoFromInput(req)
		res, err := svc.Create(r.Context(), dto)

		if errors.Is(err, entity.ErrEmailDuplicate) {
			response.Error(w, "Email already Exist", http.StatusConflict)
			return
		}
		resp := presenter.NewUserOutputFromDto(res)

		select {

		case emailChan <- worker.EmailJob{
			Email: res.Email,
			Name:  res.Name,
		}:
		default:
			log.Printf("[WARNING] Email queue is full, skipping email notification")
		}

		response.Success(w, http.StatusCreated, "User created successfully", resp)
	}
}

func FindAllUser(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := svc.FindAll(r.Context())

		if err != nil {
			response.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		outputs := make([]presenter.UserOutput, 0, len(data))

		for _, val := range data {
			outputs = append(outputs, *presenter.NewUserOutputFromDto(&val))
		}
		response.Success(w, http.StatusOK, "User fetched successfully", outputs)
	}
}

func FindByIdUser(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(entity.UserDto)

		idStr := r.PathValue("id")
		req.ID = idStr
		res, err := svc.FindById(r.Context(), req)

		if errors.Is(err, entity.ErrUserNotFound) {
			response.Error(w, "User not Found", http.StatusNotFound)
			return
		}

		resp := presenter.NewUserOutputFromDto(res)

		response.Success(w, http.StatusOK, "User fetched successfully", resp)

	}
}

func UpdateUser(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")

		req := new(entity.UserInput)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if err := validator.Validate(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		dto := entity.NewUserDtoFromInput(req)
		dto.ID = idStr

		res, err := svc.Update(r.Context(), dto)
		if errors.Is(err, entity.ErrUserNotFound) {
			response.Error(w, "User not Found", http.StatusNotFound)
			return
		}

		resp := presenter.NewUserOutputFromDto(res)

		response.Success(w, http.StatusOK, "User Updated", resp)
	}
}

func DeleteUser(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(entity.UserDto)
		idStr := r.PathValue("id")
		dto.ID = idStr

		err := svc.Delete(r.Context(), dto)
		if errors.Is(err, entity.ErrUserNotFound) {
			response.Error(w, "User not Found", http.StatusNotFound)
			return
		}
		response.SuccessNoConttent(w, http.StatusNoContent, "User Deleted")
	}
}

func CreateTransfer(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(entity.TransferInput)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := svc.Transfer(r.Context(), req.FromID, req.ToID, req.Amount); err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SuccessNoConttent(w, http.StatusOK, "Transfer successfull")
	}
}
