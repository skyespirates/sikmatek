package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type limitHandler struct {
	uc usecase.LimitUsecase
}

func NewLimitHandler(uc usecase.LimitUsecase) *limitHandler {
	return &limitHandler{
		uc: uc,
	}
}

// @Summary Limit list
// @Description returns list limit of current user
// @Tags limit
// @Produce json
// @Success 200 {object} response.Response
// @Router /limits [get]
// @Security BearerAuth
func (h *limitHandler) LimitList(w http.ResponseWriter, r *http.Request) {

	limit, err := h.uc.GetList(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{}

	resp["limit"] = limit

	utils.JSONResponse(w, "list limit", resp)

}

type Request struct {
	Requested int `json:"requested_limit"`
}

// @Summary Pengajuan limit
// @Description returns id newly created limit
// @Tags limit
// @Accept json
// @Produce json
// @Param data body Request true "Pengajuan Limit"
// @Success 201 {object} response.Response
// @Router /limits [post]
// @Security BearerAuth
func (h *limitHandler) Pengajuan(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Requested int `json:"requested_limit"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json", err)
		return
	}

	id, err := h.uc.AjukanLimit(r.Context(), payload.Requested)
	if err != nil {
		log.Printf("error: %s", err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to request limit", err)
		return
	}

	res := map[string]any{"id": id}
	response.Success(w, http.StatusCreated, "limit has been requested successfully", res)

}

// @Summary Approve limit
// @Description returns message approve status
// @Tags limit
// @Accept json
// @Produce json
// @Param data body Request true "Approve Limit"
// @Success 201 {object} response.Response
// @Router /limits/pengajuan-limit/:limit_id/approve [patch]
// @Security BearerAuth
func (h *limitHandler) Approve(w http.ResponseWriter, r *http.Request) {

	var payload entity.UpdateLimitPayload

	ps := httprouter.ParamsFromContext(r.Context())
	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request, limit id must be a number", err)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "APPROVED"

	err = h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to approve limit", err)
		return
	}

	response.Success(w, http.StatusOK, "limit has been approved successfully", nil)

}

// @Summary Reject limit
// @Description returns message reject status
// @Tags limit
// @Accept json
// @Produce json
// @Param data body Request true "Reject Limit"
// @Success 201 {object} response.Response
// @Router /limits [patch]
// @Security BearerAuth
func (h *limitHandler) Reject(w http.ResponseWriter, r *http.Request) {

	var payload entity.UpdateLimitPayload

	ps := httprouter.ParamsFromContext(r.Context())

	id := ps.ByName("limit_id")

	limit_id, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request, limit id must be a number", err)
		return
	}

	payload.LimitId = limit_id
	payload.Action = "REJECTED"

	err = h.uc.TindakLanjut(r.Context(), payload)
	if err != nil {
		log.Printf("error: %s", err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to reject limit", err)
		return
	}

	response.Success(w, http.StatusOK, "limit has been rejected successfully", nil)

}

func (h *limitHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}

func (h *limitHandler) ListApproved(w http.ResponseWriter, r *http.Request) {

	limits, err := h.uc.ListLimitAktif(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get list of approved limit", err)
		return
	}

	resp := map[string]any{"limits": limits}
	response.Success(w, http.StatusOK, "list of approved limit", resp)

}
