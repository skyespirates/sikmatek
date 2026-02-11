package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
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

func (h installmentHandler) PayInstallment(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())

	rawId := ps.ByName("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}

	err = h.uc.PayInstallment(r.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrDuplicatePayment) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		msg := fmt.Sprintf("internal server error: %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "cicilan berhasil dibayar", struct{}{})

}
