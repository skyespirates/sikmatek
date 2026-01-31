package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/delivery/http/handler"
	"github.com/skyespirates/sikmatek/internal/infra/pgsql"
	"github.com/skyespirates/sikmatek/internal/usecase"
	"github.com/skyespirates/sikmatek/internal/utils"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	taskHandler := handler.NewTaskHandler(usecase.NewTaskUsecase(pgsql.NewTaskRepository(app.db)))
	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(pgsql.NewUserRepository(app.db)))

	router.HandlerFunc(http.MethodGet, "/", index)
	router.HandlerFunc(http.MethodGet, "/healthcheck", healthcheck)

	router.HandlerFunc(http.MethodPost, "/v1/auth/register", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/auth/login", userHandler.Login)

	router.HandlerFunc(http.MethodGet, "/v1/tasks", app.authenticate(taskHandler.GetAll))
	router.HandlerFunc(http.MethodGet, "/v1/tasks/:id", taskHandler.GetById)
	router.HandlerFunc(http.MethodPost, "/v1/tasks", app.authenticate(taskHandler.Create))
	router.HandlerFunc(http.MethodPut, "/v1/tasks/:id", app.authenticate(taskHandler.Update))
	router.HandlerFunc(http.MethodDelete, "/v1/tasks/:id", taskHandler.Delete)

	router.HandlerFunc(http.MethodPost, "/generate-key", func(w http.ResponseWriter, r *http.Request) {
		res := make(map[string]string)

		res["key"] = utils.GenerateKey()

		json.NewEncoder(w).Encode(res)
	})

	router.HandlerFunc(http.MethodPost, "/encrypt", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Key  string `json:"key"`
			Text string `json:"text"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if input.Key == "" || input.Text == "" {
			http.Error(w, "bad request, key and text are required", http.StatusBadRequest)
			return
		}

		result := utils.Encrypt(input.Key, input.Text)

		res := make(map[string]string)
		res["encrypted"] = result
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "error on json encoder", http.StatusInternalServerError)
		}
	})

	router.HandlerFunc(http.MethodPost, "/decrypt", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Key  string `json:"key"`
			Text string `json:"text"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if input.Key == "" || input.Text == "" {
			http.Error(w, "bad request, key and encrypted text are required", http.StatusBadRequest)
			return
		}

		decoded := utils.Decrypt(input.Key, input.Text)
		res := make(map[string]string)
		res["decrypted"] = decoded

		json.NewEncoder(w).Encode(res)

	})

	return app.loggerMiddleware(app.corsMiddleware(router))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Skyes! 😎"))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All iz well 👌"))
}
