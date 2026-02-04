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

	// repositories
	userRepo := mysql.NewUserRepository(app.db)
	consumerRepo := mysql.NewConsumerRepository()
	limitRepo := mysql.NewLimitRepository()
	contractRepo := mysql.NewContractRepository()
	productRepo := mysql.NewProductRepository()

	// usecases
	userUC := usecase.NewUserUsecase(app.db, userRepo, consumerRepo)
	consumerUC := usecase.NewConsumerUsecase(app.db, consumerRepo)
	limitUC := usecase.NewLimitUsecase(app.db, limitRepo)
	contractUC := usecase.NewContractUsecase(app.db, contractRepo, limitRepo, productRepo)
	productUC := usecase.NewProductUsecase(app.db, productRepo)

	// handlers
	userHandler := handler.NewUserHandler(userUC)
	consumerHandler := handler.NewConsumerHandler(consumerUC)
	limitHandler := handler.NewLimitHandler(limitUC)
	contractHandler := handler.NewContractHandler(contractUC)
	productHandler := handler.NewProductHandler(productUC)

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
	router.HandlerFunc(http.MethodPost, "/v1/pengajuan-limit", app.authenticate(limitHandler.Pengajuan))
	router.HandlerFunc(http.MethodPatch, "/v1/pengajuan-limit/:limit_id/approve", limitHandler.Approve)
	router.HandlerFunc(http.MethodPatch, "/v1/pengajuan-limit/:limit_id/reject", limitHandler.Reject)

	// products service
	router.HandlerFunc(http.MethodGet, "/v1/products", productHandler.List)
	router.HandlerFunc(http.MethodPost, "/v1/products", productHandler.Create)

	// contracts service
	router.HandlerFunc(http.MethodPost, "/v1/kontrak", app.authenticate(contractHandler.BuatKontrak))
	router.HandlerFunc(http.MethodGet, "/v1/kontrak", app.authenticate(contractHandler.ListKontrak))
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
