package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "get profile", info)

}

func (h *consumerHandler) CompleteConsumerInfo(w http.ResponseWriter, r *http.Request) {
	var payload entity.UpdateConsumerPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.uc.CompleteInfo(r.Context(), payload)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, mysql.ErrDuplicateNik) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "profile has been updated", nil)

}
func (h *consumerHandler) UploadKtp(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := h.c.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "KTP has been uploaded", nil)
}
func (h *consumerHandler) UploadSelfie(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := h.c.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, "Selfie has been uploaded", nil)

}

func (h *consumerHandler) VerifyConsumer(w http.ResponseWriter, r *http.Request) {

	ps := httprouter.ParamsFromContext(r.Context())

	consumerID := ps.ByName("consumer_id")

	err := h.uc.Verify(r.Context(), consumerID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "consumer is verifed successfully")

}

func (h *consumerHandler) CheckLimit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}
