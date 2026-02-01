package handler

import "net/http"

type consumerHandler struct{}

func NewConsumerHandler() *consumerHandler {
	return &consumerHandler{}
}

func (h *consumerHandler) CreateConsumerInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create consumer info"))
}
func (h *consumerHandler) UploadKtp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("upload ktp"))
}
func (h *consumerHandler) UploadSelfie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("upload selfie"))
}
func (h *consumerHandler) CheckLimit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check limit"))
}
