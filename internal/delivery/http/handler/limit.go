package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type limitHandler struct {
	uc usecase.LimitUsecase
}

func NewLimitHandler(uc usecase.LimitUsecase) *limitHandler {
	return &limitHandler{
		uc: uc,
	}
}

func (h *limitHandler) LimitList(w http.ResponseWriter, r *http.Request) {

	limit, err := h.uc.GetList(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}

	resp["limit"] = limit

	utils.JSONResponse(w, "list limit", resp)

}

func (h *limitHandler) Pengajuan(w http.ResponseWriter, r *http.Request) {

	var payload entity.CreateLimitPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := h.uc.AjukanLimit(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	location := fmt.Sprintf("/v1/limits/%d", id)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", location)

	fmt.Fprintf(w, `{"id": %d}`, id)

}

func (h *limitHandler) Approve(w http.ResponseWriter, r *http.Request) {

	var payload entity.UpdateLimitPayload

	ps := httprouter.ParamsFromContext(r.Context())
	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad request, limit id must be a number", http.StatusBadRequest)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "APPROVED"

	err = h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["message"] = "limit approved"

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}

}

func (h *limitHandler) Reject(w http.ResponseWriter, r *http.Request) {

	var payload entity.UpdateLimitPayload

	ps := httprouter.ParamsFromContext(r.Context())

	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad request, limit id must be a number", http.StatusBadRequest)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "REJECTED"

	err = h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["message"] = "limit rejected'"

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}

}

func (h *limitHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}
