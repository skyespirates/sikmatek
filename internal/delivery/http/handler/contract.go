package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
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
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nomor_kontrak, err := h.uc.Create(r.Context(), payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resource := fmt.Sprintf("/v1/kontrak/%s", nomor_kontrak)
	w.Header().Set("Location", resource)
	w.WriteHeader(http.StatusCreated)
}

func (h *contractHandler) ListKontrak(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Kontrak"))
}

func (h *contractHandler) QuoteKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")
	err := h.uc.GenerateQuote(r.Context(), nomor_kontrak)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "contract %s has quoted", nomor_kontrak)

}

func (h *contractHandler) ConfirmKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Confirm(r.Context(), nomor_kontrak)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "contract %s has confirmed", nomor_kontrak)
	w.WriteHeader(http.StatusNoContent)

}

func (h *contractHandler) CancelKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Confirm(r.Context(), nomor_kontrak)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "contract %s has cancelled", nomor_kontrak)
	w.WriteHeader(http.StatusNoContent)

}

func (h *contractHandler) ActivateKontrak(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())
	nomor_kontrak := ps.ByName("nomor_kontrak")

	err := h.uc.Confirm(r.Context(), nomor_kontrak)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "contract %s has activated", nomor_kontrak)
	w.WriteHeader(http.StatusNoContent)

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
