package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type productHandler struct {
	uc usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *productHandler {
	return &productHandler{
		uc: uc,
	}
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {

	var payload entity.CreateProductPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := h.uc.Create(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v1/products/%d", id))
	w.WriteHeader(http.StatusCreated)

}

func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.uc.GetList(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}
	resp["message"] = "list of products"
	resp["products"] = products

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error encoding", http.StatusInternalServerError)
	}
}
