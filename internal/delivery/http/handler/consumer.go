package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/response"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type consumerHandler struct {
	uc usecase.ConsumerUsecase
	c  *cloudinary.Cloudinary
}

func NewConsumerHandler(uc usecase.ConsumerUsecase, c *cloudinary.Cloudinary) *consumerHandler {
	return &consumerHandler{
		uc: uc,
		c:  c,
	}
}

func (h *consumerHandler) GetConsumerInfo(w http.ResponseWriter, r *http.Request) {

	info, err := h.uc.GetInfo(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get customer info", err)
		return
	}

	response.Success(w, http.StatusOK, "get customer profile successfully", info)

}

func (h *consumerHandler) CompleteConsumerInfo(w http.ResponseWriter, r *http.Request) {
	var payload entity.UpdateConsumerPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusBadRequest, "failed to decode json", err)
		return
	}

	err = h.uc.CompleteInfo(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, mysql.ErrDuplicateNik) {
			response.Error(w, http.StatusConflict, "nik already used", err)
			return
		}
		response.Error(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	response.Success(w, http.StatusOK, "profile has been updated", nil)

}
func (h *consumerHandler) UploadKtp(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "file size exceeded", err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "error retrieving the file", err)
		return
	}
	defer file.Close()

	result, err := h.c.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to upload ktp", err)
		return
	}

	// safeName := strings.ReplaceAll(handler.Filename, " ", "_")
	// filePath := filepath.Join("static", "uploads", "ktp", safeName)
	// dst, err := os.Create(filePath)
	// if err != nil {
	// 	log.Printf("error: %s", err.Error())
	// 	http.Error(w, "error on create destination", http.StatusInternalServerError)
	// 	return
	// }
	// defer dst.Close()

	// _, err = io.Copy(dst, file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	claims := utils.ContextGetUser(r.Context())

	err = h.uc.SetKtp(r.Context(), claims.ConsumerId, result.URL)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to set ktp", err)
		return
	}

	response.Success(w, http.StatusOK, "ktp has been updated successfully", nil)

}
func (h *consumerHandler) UploadSelfie(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "file size exceeded", err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "error retrieving the file", err)
		return
	}
	defer file.Close()

	result, err := h.c.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to upload ktp", err)
		return
	}

	// safeName := strings.ReplaceAll(handler.Filename, " ", "_")
	// filePath := filepath.Join("static", "uploads", "selfie", safeName)
	// dst, err := os.Create(filePath)
	// if err != nil {
	// 	log.Printf("error: %s", err.Error())
	// 	http.Error(w, "error on create destination", http.StatusInternalServerError)
	// 	return
	// }
	// defer dst.Close()

	// _, err = io.Copy(dst, file)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	claims := utils.ContextGetUser(r.Context())

	err = h.uc.SetSelfie(r.Context(), claims.ConsumerId, result.URL)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusInternalServerError, "failed to set selfie", err)
		return
	}

	response.Success(w, http.StatusOK, "selfie has been updated successfully", nil)

}

func (h *consumerHandler) VerifyConsumer(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())

	consumerID := ps.ByName("consumer_id")

	err := h.uc.Verify(r.Context(), consumerID)
	if err != nil {
		log.Println(err.Error())
		response.Error(w, http.StatusBadRequest, "failed to verify consumer", err)
		return
	}

	response.Success(w, http.StatusOK, "consumer has been verified successfully", nil)

}

func (h *consumerHandler) CheckLimit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}
