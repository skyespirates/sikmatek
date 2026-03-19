package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/response"
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
		response.Error(w, http.StatusInternalServerError, "failed to generate installment", err)
		return
	}

	response.Success(w, http.StatusOK, "installment has generated successfully", nil)
}

func (h installmentHandler) PayInstallment(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())

	rawId := ps.ByName("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "id must be a number", err)
		return
	}

	err = h.uc.PayInstallment(r.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrDuplicatePayment) {
			response.Error(w, http.StatusConflict, "conflict duplicate", err)
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to pay installment", err)
		return
	}

	response.Success(w, http.StatusOK, "installment has payed successfully", nil)

}

func (h installmentHandler) ListInstallment(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())

	nomor_kontrak := ps.ByName("nomor_kontrak")

	installements, err := h.uc.ListInstallment(r.Context(), nomor_kontrak)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to get installment list", err)
		return
	}

	res := map[string]any{"installments": installements}

	response.Success(w, http.StatusOK, "get installment list", res)

}
