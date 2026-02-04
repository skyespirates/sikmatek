package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/delivery/http/handler"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(app.db, mysql.NewUserRepository(app.db), mysql.NewConsumerRepository()))
	consumerHandler := handler.NewConsumerHandler(usecase.NewConsumerUsecase(app.db, mysql.NewConsumerRepository()))
	limitHandler := handler.NewLimitHandler(usecase.NewLimitUsecase(app.db, mysql.NewLimitRepository()))
	contractHandler := handler.NewContractHandler(usecase.NewContractUsecase(mysql.NewContractRepository()))

	// serve static files, foto ktp dan selfie bisa diakses di sini
	router.ServeFiles("/assets/*filepath", http.Dir("client/dist/assets"))
	router.ServeFiles("/uploads/*filepath", http.Dir("static/uploads"))

	router.HandlerFunc(http.MethodGet, "/", index)
	router.HandlerFunc(http.MethodGet, "/healthcheck", healthcheck)

	// auth service
	router.HandlerFunc(http.MethodPost, "/v1/auth/register", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/auth/login", userHandler.Login)

	// consumers service
	router.HandlerFunc(http.MethodPut, "/v1/consumers", app.authenticate(consumerHandler.CompleteConsumerInfo))
	router.HandlerFunc(http.MethodPut, "/v1/consumers/upload-ktp", app.authenticate(consumerHandler.UploadKtp))
	router.HandlerFunc(http.MethodPut, "/v1/consumers/upload-selfie", app.authenticate(consumerHandler.UploadSelfie))
	router.HandlerFunc(http.MethodPatch, "/v1/consumers/:consumer_id/verify", app.authenticate(consumerHandler.VerifyConsumer))
	router.HandlerFunc(http.MethodGet, "/v1/consumers/limits/:limit_id/sisa-limit", consumerHandler.CheckLimit)

	// limits service
	router.HandlerFunc(http.MethodGet, "/v1/limits", app.authenticate(limitHandler.LimitList))
	router.HandlerFunc(http.MethodPost, "/v1/pengajuan-limit", limitHandler.Pengajuan)
	router.HandlerFunc(http.MethodPatch, "/v1/pengajuan-limit/:limit_id/approve", limitHandler.Approve)
	router.HandlerFunc(http.MethodPatch, "/v1/pengajuan-limit/:limit_id/reject", limitHandler.Reject)

	// products service
	router.HandlerFunc(http.MethodGet, "/v1/products", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/products", userHandler.Register)

	// contracts service
	router.HandlerFunc(http.MethodPost, "/v1/kontrak", contractHandler.BuatKontrak)
	router.HandlerFunc(http.MethodGet, "/v1/kontrak", contractHandler.ListKontrak)
	router.HandlerFunc(http.MethodPatch, "/v1/kontrak/:nomor_kontrak/quote", contractHandler.QuoteKontrak)
	router.HandlerFunc(http.MethodPatch, "/v1/kontrak/:nomor_kontrak/confirm", contractHandler.ConfirmKontrak)
	router.HandlerFunc(http.MethodPatch, "/v1/kontrak/:nomor_kontrak/cancel", contractHandler.CancelKontrak)
	router.HandlerFunc(http.MethodPatch, "/v1/kontrak/:nomor_kontrak/activate", contractHandler.ActivateKontrak)
	router.HandlerFunc(http.MethodPatch, "/v1/kontrak/:nomor_kontrak/cicilan", contractHandler.CicilKontrak)
	router.HandlerFunc(http.MethodGet, "/v1/kontrak/:nomor_kontrak", contractHandler.DetailKontrak)
	router.HandlerFunc(http.MethodGet, "/v1/kontrak/:nomor_kontrak/installments", contractHandler.DaftarCicilan)

	router.HandlerFunc(http.MethodPost, "/v1/cicilan/:id/bayar", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprintln(w, "No name provided")
			return
		}

		w.Write([]byte(name))
	})

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./client/dist/index.html")
	})
	return app.loggerMiddleware(app.corsMiddleware(router))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/dist/index.html")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All iz well 👌"))
}
