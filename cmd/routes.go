package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/delivery/http/handler"
	"github.com/skyespirates/sikmatek/internal/infra/mysql"
	"github.com/skyespirates/sikmatek/internal/infra/pgsql"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(pgsql.NewUserRepository(app.db)))
	tenorHandler := handler.NewTenorHandler(usecase.NewTenorUsecase(mysql.NewTenorRepository(app.db)))

	router.ServeFiles("/assets/*filepath", http.Dir("client/dist/assets"))

	router.HandlerFunc(http.MethodGet, "/", index)
	router.HandlerFunc(http.MethodGet, "/healthcheck", healthcheck)

	router.HandlerFunc(http.MethodPost, "/v1/auth/register", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/consumers", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/tenors", tenorHandler.GetList)
	router.HandlerFunc(http.MethodPost, "/v1/tenors", tenorHandler.CreateTenor)
	router.HandlerFunc(http.MethodPost, "/v1/pengajuan-limit", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/pengajuan-limit/:id/approve", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/pengajuan-limit/:id/reject", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/produk", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/kontrak", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/kontrak/:nomor_kontrak", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/kontrak/:nomor_kontrak/cicilan", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/cicilan/:id/bayar", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/konsumen/:id/sisa-limit", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/credit-limits", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/transactions", userHandler.Register)
	router.HandlerFunc(http.MethodGet, "/v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprintln(w, "No name provided")
			return
		}

		w.Write([]byte(name))
	})

	router.HandlerFunc(http.MethodPost, "/v1/uploads", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "file too large", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		dst, err := os.Create("./static/uploads/" + handler.Filename)
		if err != nil {
			app.logger.LogInfo(r, err.Error())
			http.Error(w, "error on create destination", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "file uploaded successfully")

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
