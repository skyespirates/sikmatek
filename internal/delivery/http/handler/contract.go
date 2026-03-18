package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type contractHandler struct {
	uc usecase.ContractUsecase
}

func NewContractHandler(uc usecase.ContractUsecase) *contractHandler {
	return &contractHandler{
		uc: uc,
	}
}

func (h *contractHandler) BuatKontrak(w http.ResponseWriter, r *http.Request) {
	var payload entity.CreateContractPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusBadRequest, "bad request", err)
		return
	}

	nomor_kontrak, err := h.uc.Create(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}

	resp["nomor_kontrak"] = nomor_kontrak

	response.Success(w, http.StatusCreated, "contract has been created successfully", resp)
}

func (h *contractHandler) ListKontrak(w http.ResponseWriter, r *http.Request) {

	contract, err := h.uc.DaftarKontrak(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}
	resp["kontrak"] = contract

	response.Success(w, http.StatusOK, "list of contract", resp)

}

func (h *contractHandler) QuoteKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")
	err := h.uc.GenerateQuote(r.Context(), nomor_kontrak)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}
	resp["nomor_kontrak"] = nomor_kontrak

	response.Success(w, http.StatusOK, "contract has been quoted successfully", resp)

}

func (h *contractHandler) ConfirmKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Confirm(r.Context(), nomor_kontrak)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}
	resp["nomor_kontrak"] = nomor_kontrak

	response.Success(w, http.StatusOK, "contract has been confirmed successfully", resp)
}

func (h *contractHandler) CancelKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Cancel(r.Context(), nomor_kontrak)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}
	resp["nomor_kontrak"] = nomor_kontrak

	response.Success(w, http.StatusOK, "contract has been cancelled successfully", resp)

}

func (h *contractHandler) ActivateKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Activate(r.Context(), nomor_kontrak)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	resp := map[string]any{}
	resp["nomor_kontrak"] = nomor_kontrak

	response.Success(w, http.StatusOK, "contract has been activated successfully", resp)

}

func (h *contractHandler) CicilKontrak(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("konsument mencicil kontrak"))
}

func (h *contractHandler) DetailKontrak(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("detail kontrak"))
}

func (h *contractHandler) DaftarCicilan(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menampilkan daftar cicilan suatu kontrak"))
}
