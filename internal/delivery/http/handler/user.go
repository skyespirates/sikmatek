package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/validator"
)

type userHandler struct {
	uc usecase.UserUsecase
	v  *validator.Validator
}

func NewUserHandler(uc usecase.UserUsecase, v *validator.Validator) *userHandler {
	return &userHandler{uc: uc, v: v}
}

// @Summary Register new user
// @Description returns jwt token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body entity.RegisterPayload true "Register Data"
// @Success 200 {object} response.Response
// @Router /auth/register [post]
func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {

	var payload entity.RegisterPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.v.Validate(payload)
	if err != nil {
		errors := validator.FormatErrors(err)
		response.Error(w, http.StatusBadRequest, "validation error", errors)
		return
	}

	token, err := h.uc.Register(r.Context(), &payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		if errors.Is(err, mysql.ErrDuplicate) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	resp := map[string]any{}
	resp["token"] = token

	response.Success(w, http.StatusCreated, "registered successfully", resp)

}

// @Summary Login user
// @Description returns jwt token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body entity.LoginPayload true "Login Data"
// @Success 200 {object} response.Response
// @Router /auth/login [post]
func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var payload entity.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.v.Validate(payload)
	if err != nil {
		errors := validator.FormatErrors(err)
		response.Error(w, http.StatusBadRequest, "validation error", errors)
		return
	}

	token, err := h.uc.Login(r.Context(), &payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		if errors.Is(err, usecase.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := make(map[string]interface{})
	res["token"] = token

	response.Success(w, http.StatusOK, "login successfully", res)
}
