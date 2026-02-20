package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/delivery/http/handler"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
)

//go:embed dist
var embeddedFiles embed.FS

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// repositories
	userRepo := mysql.NewUserRepository(app.db)
	consumerRepo := mysql.NewConsumerRepository()
	limitRepo := mysql.NewLimitRepository()
	contractRepo := mysql.NewContractRepository()
	productRepo := mysql.NewProductRepository()
	installmentRepo := mysql.NewInstallmentRepository()
	limitUsageRepo := mysql.NewLimitUsageRepository()

	// usecases
	userUC := usecase.NewUserUsecase(app.db, userRepo, consumerRepo)
	consumerUC := usecase.NewConsumerUsecase(app.db, consumerRepo)
	limitUC := usecase.NewLimitUsecase(app.db, limitRepo)
	contractUC := usecase.NewContractUsecase(app.db, contractRepo, limitRepo, productRepo, installmentRepo)
	productUC := usecase.NewProductUsecase(app.db, productRepo)
	installmentUC := usecase.NewInstallmentUsecase(app.db, installmentRepo, contractRepo, limitUsageRepo)
	dashboardUC := usecase.NewDashboardUsecase(app.db, consumerRepo, limitRepo, contractRepo, productRepo)

	// handlers
	userHandler := handler.NewUserHandler(userUC)
	consumerHandler := handler.NewConsumerHandler(consumerUC)
	limitHandler := handler.NewLimitHandler(limitUC)
	contractHandler := handler.NewContractHandler(contractUC)
	productHandler := handler.NewProductHandler(productUC)
	installmentHandler := handler.NewInstallmentHandler(installmentUC)
	dashboardHandler := handler.NewDashboardHandler(dashboardUC)

	// serve static files, foto ktp dan selfie bisa diakses di sini
	// router.ServeFiles("/assets/*filepath", http.Dir("client/dist/assets"))
	router.ServeFiles("/uploads/*filepath", http.Dir("static/uploads"))

	// distDir := "./dist"

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	// Try to find requested file
	// 	path := filepath.Join(distDir, r.URL.Path)
	// 	_, err := os.Stat(path)

	// 	// If file does not exist → serve index.html (React Router)
	// 	if os.IsNotExist(err) {
	// 		http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
	// 		return
	// 	}

	// 	fileServer.ServeHTTP(w, r)
	// })
	router.HandlerFunc(http.MethodGet, "/api/healthcheck", healthcheck)

	// auth service
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/register", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/login", userHandler.Login)

	// consumers service
	router.HandlerFunc(http.MethodGet, "/api/v1/consumers", app.authenticate(app.authorize(utils.Roles["admin"], utils.Roles["consumer"])(consumerHandler.GetConsumerInfo)))
	router.HandlerFunc(http.MethodPut, "/api/v1/consumers", app.authenticate(app.authorize(utils.Roles["consumer"])(consumerHandler.CompleteConsumerInfo)))
	router.HandlerFunc(http.MethodPut, "/api/v1/consumers/upload-ktp", app.authenticate(app.authorize(utils.Roles["consumer"])(consumerHandler.UploadKtp)))
	router.HandlerFunc(http.MethodPut, "/api/v1/consumers/upload-selfie", app.authenticate(app.authorize(utils.Roles["consumer"])(consumerHandler.UploadSelfie)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/consumers/:consumer_id/verify", app.authenticate(app.authorize(utils.Roles["admin"])(consumerHandler.VerifyConsumer)))
	router.HandlerFunc(http.MethodGet, "/api/v1/consumers/limits/:limit_id/sisa-limit", consumerHandler.CheckLimit)

	// limits service
	router.HandlerFunc(http.MethodGet, "/api/v1/limits", app.authenticate(app.authorize(utils.Roles["admin"], utils.Roles["consumer"])(limitHandler.LimitList)))
	router.HandlerFunc(http.MethodPost, "/api/v1/pengajuan-limit", app.authenticate(app.authorize(utils.Roles["consumer"])(limitHandler.Pengajuan)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/pengajuan-limit/:limit_id/approve", app.authenticate(app.authorize(utils.Roles["admin"])(limitHandler.Approve)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/pengajuan-limit/:limit_id/reject", app.authenticate(app.authorize(utils.Roles["admin"])(limitHandler.Reject)))
	router.HandlerFunc(http.MethodGet, "/api/v1/limits/approved", app.authenticate(app.authorize(utils.Roles["admin"], utils.Roles["consumer"])(limitHandler.ListApproved)))

	// products service
	router.HandlerFunc(http.MethodGet, "/api/v1/products", app.authenticate(app.authorize(utils.Roles["admin"], utils.Roles["consumer"])(productHandler.List)))
	router.HandlerFunc(http.MethodPost, "/api/v1/products", app.authenticate(app.authorize(utils.Roles["admin"])(productHandler.Create)))

	// contracts service
	router.HandlerFunc(http.MethodPost, "/api/v1/kontrak", app.authenticate(app.authorize(utils.Roles["consumer"])(contractHandler.BuatKontrak)))
	router.HandlerFunc(http.MethodGet, "/api/v1/kontrak", app.authenticate(contractHandler.ListKontrak))
	router.HandlerFunc(http.MethodPatch, "/api/v1/kontrak/:nomor_kontrak/quote", app.authenticate(app.authorize(utils.Roles["admin"])(contractHandler.QuoteKontrak)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/kontrak/:nomor_kontrak/confirm", app.authenticate(app.authorize(utils.Roles["consumer"])(contractHandler.ConfirmKontrak)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/kontrak/:nomor_kontrak/cancel", app.authenticate(app.authorize(utils.Roles["consumer"])(contractHandler.CancelKontrak)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/kontrak/:nomor_kontrak/activate", app.authenticate(app.authorize(utils.Roles["admin"])(contractHandler.ActivateKontrak)))
	router.HandlerFunc(http.MethodPatch, "/api/v1/kontrak/:nomor_kontrak/cicilan", app.authenticate(app.authorize(utils.Roles["consumer"])(contractHandler.CicilKontrak)))
	router.HandlerFunc(http.MethodGet, "/api/v1/kontrak/:nomor_kontrak", app.authenticate(contractHandler.DetailKontrak))

	// installment service
	router.HandlerFunc(http.MethodPut, "/api/v1/installments/:id", app.authenticate(installmentHandler.PayInstallment))

	// dashboard service
	router.HandlerFunc(http.MethodGet, "/api/v1/dashboard/consumer", app.authenticate(app.authorize(utils.Roles["consumer"])(dashboardHandler.GetConsumerDashboard)))

	router.HandlerFunc(http.MethodPost, "/api/v1/cicilan/:id/bayar", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/api/v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprintln(w, "No name provided")
			return
		}

		w.Write([]byte(name))
	})

	router.HandlerFunc(http.MethodPost, "/api/v1/installments/:nomor_kontrak/generate", app.authenticate(app.authorize(utils.Roles["admin"])(installmentHandler.GenerateInstallment)))

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "endpoint not found"}`))
			return
		}

		serveReactApp(w, r)
	})

	return app.loggerMiddleware(router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All iz well 👌"))
}

func serveReactApp(w http.ResponseWriter, r *http.Request) {
	distFS, _ := fs.Sub(embeddedFiles, "dist")
	fileServer := http.FileServer(http.FS(distFS))

	filePath := strings.TrimPrefix(r.URL.Path, "/")
	f, err := distFS.Open(filePath)
	if err == nil {
		f.Close()
		fileServer.ServeHTTP(w, r)
		return
	}

	index, err := embeddedFiles.ReadFile("dist/index.html")
	if err != nil {
		http.Error(w, "index.html not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(index)
}
