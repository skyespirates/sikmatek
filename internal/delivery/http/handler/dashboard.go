package handler

import (
	"net/http"

	"github.com/skyespirates/sikmatek/internal/response"
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
		response.Error(w, http.StatusInternalServerError, "failed to get dashboard data", err)
		return
	}

	response.Success(w, http.StatusOK, "get dashboard data successfully", data)
}
