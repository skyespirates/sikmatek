package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/usecase"
)

type tenorHandler struct {
	uc usecase.TenorUsecase
}

func NewTenorHandler(uc usecase.TenorUsecase) *tenorHandler {
	return &tenorHandler{
		uc: uc,
	}
}

func (h *tenorHandler) CreateTenor(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DurasiBulan int `json:"durasi_bulan"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := h.uc.Create(r.Context(), input.DurasiBulan)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"message": "tenor created successfully",
		"id":      id,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}
}

func (h *tenorHandler) GetList(w http.ResponseWriter, r *http.Request) {
	tenors, err := h.uc.GetAll(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := make(map[string]any)
	res["tenors"] = tenors

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error on encoding", http.StatusInternalServerError)
	}
}
