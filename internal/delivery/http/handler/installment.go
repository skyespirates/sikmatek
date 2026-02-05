package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type installmentHandler struct {
	uc usecase.InstallmentUsecase
}

func NewInstallmentHandler(uc usecase.InstallmentUsecase) *installmentHandler {
	return &installmentHandler{
		uc: uc,
	}
}

func (h *installmentHandler) GenerateInstallment(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.GenerateInstallment(r.Context(), nomor_kontrak)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "installment has generated successfully")
	w.WriteHeader(http.StatusCreated)
}
