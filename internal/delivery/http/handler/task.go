package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/usecase"
)

type taskHandler struct {
	uc usecase.TaskUsecase
}

func NewTaskHandler(uc usecase.TaskUsecase) *taskHandler {
	return &taskHandler{uc}
}

func (th *taskHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	tasks, err := th.uc.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"data": tasks,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (th *taskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	task, err := th.uc.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"data": task,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (th *taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := th.uc.Create(r.Context(), req.Title)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]any)
	response["message"] = "task created successfully"
	response["task"] = task
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (th *taskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	todoId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	deletedId, err := th.uc.Delete(r.Context(), todoId)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := make(map[string]any)
	response["message"] = "todo deleted successfully"
	response["id"] = deletedId
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error: %s", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (th *taskHandler) Update(w http.ResponseWriter, r *http.Request) {
	updatedTask, err := th.uc.Update(r.Context(), r)
	res := make(map[string]any)
	if err != nil {
		res["status"] = "error"
		res["message"] = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	res["status"] = "success"
	res["message"] = "updated successfully"
	res["task"] = updatedTask
	json.NewEncoder(w).Encode(res)
}
