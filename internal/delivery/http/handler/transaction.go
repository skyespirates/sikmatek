package handler

import "net/http"

type transactionHandler struct{}

func NewTransactionHandler() *transactionHandler {
	return &transactionHandler{}
}

func (h *transactionHandler) BuatKontrak(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buat kontrak"))
}

func (h *transactionHandler) DetailKontrak(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("detail kontrak"))
}

func (h *transactionHandler) CicilanKontrak(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cicialan kontrak"))
}
