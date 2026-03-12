package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/validator"
)

type productHandler struct {
	uc usecase.ProductUsecase
	v  *validator.Validator
}

func NewProductHandler(uc usecase.ProductUsecase, v *validator.Validator) *productHandler {
	return &productHandler{
		uc: uc,
		v:  v,
	}
}

// @Summary Create a product
// @Description returns id of newly created product
// @Tags product
// @Accept json
// @Produce json
// @Param data body entity.CreateProductPayload true "Create Product"
// @Success 200 {object} response.Response
// @Router /products [post]
// @Security BearerAuth
func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {

	var payload entity.CreateProductPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err.Error())
		errs := validator.FormatErrors(err)
		response.Error(w, http.StatusBadRequest, "decoding error", errs)
		return
	}

	err = h.v.Validate(payload)
	if err != nil {
		log.Println(err.Error())
		errors := validator.FormatErrors(err)
		response.Error(w, http.StatusBadRequest, "validation error", errors)
		return
	}

	id, err := h.uc.Create(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{"id_product": id}

	response.Success(w, http.StatusCreated, "product created successfully", resp)

}

// @Summary Get list of products
// @Description returns list of products
// @Tags product
// @Produce json
// @Success 200 {object} response.Response
// @Router /products [get]
// @Security BearerAuth
func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.uc.GetList(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, http.StatusOK, "list products", products)
}
