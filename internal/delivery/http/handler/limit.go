package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type limitHandler struct {
	uc usecase.LimitUsecase
}

func NewLimitHandler(uc usecase.LimitUsecase) *limitHandler {
	return &limitHandler{
		uc: uc,
	}
}

func (h *limitHandler) Pengajuan(w http.ResponseWriter, r *http.Request) {
	var payload entity.CreateLimitPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	limit, err := h.uc.AjukanLimit(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["limit"] = limit

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}
}

func (h *limitHandler) Approve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var payload entity.UpdateLimitPayload

	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad request, limit id must be a number", http.StatusBadRequest)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "APPROVED"

	limit, err := h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["message"] = "limit approved"
	resp["limit"] = limit

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}
}
func (h *limitHandler) Reject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var payload entity.UpdateLimitPayload

	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad request, limit id must be a number", http.StatusBadRequest)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "REJECTED"

	limit, err := h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["message"] = "limit rejected'"
	resp["limit"] = limit

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}
}

func (h *limitHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}
