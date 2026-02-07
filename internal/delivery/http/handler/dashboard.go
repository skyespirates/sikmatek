package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/usecase"
)

type dashboardHandler struct {
	uc usecase.DashboardUsecase
}

func NewDashboardHandler(uc usecase.DashboardUsecase) *dashboardHandler {
	return &dashboardHandler{
		uc: uc,
	}
}

func (h *dashboardHandler) GetConsumerDashboard(w http.ResponseWriter, r *http.Request) {
	data, err := h.uc.GetConsumerDashboardData(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
