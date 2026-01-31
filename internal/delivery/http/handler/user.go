package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/infra/pgsql"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type userHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *userHandler {
	return &userHandler{uc: uc}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload entity.RegisterPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	u, err := h.uc.Register(r.Context(), &payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		if errors.Is(err, pgsql.DuplicateErr) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	jsonByte, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.Write(jsonByte)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload entity.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	token, err := h.uc.Login(r.Context(), &payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		if errors.Is(err, pgsql.ErrNotFound) {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := make(map[string]string)
	res["message"] = "login successfully"
	res["token"] = token

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
