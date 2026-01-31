package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type TaskUsecase interface {
	GetAll(context.Context) ([]*entity.Task, error)
	GetById(context.Context, string) (*entity.Task, error)
	Create(context.Context, string) (*entity.Task, error)
	Delete(context.Context, int) (int, error)
	Update(context.Context, *http.Request) (*entity.Task, error)
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repo: repo,
	}
}

func (tu *taskUsecase) GetAll(ctx context.Context) ([]*entity.Task, error) {
	tasks, err := tu.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tu *taskUsecase) GetById(ctx context.Context, id string) (*entity.Task, error) {
	task_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	task, err := tu.repo.GetById(ctx, task_id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tu *taskUsecase) Create(ctx context.Context, title string) (*entity.Task, error) {
	task, err := tu.repo.Create(ctx, title)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tu *taskUsecase) Delete(ctx context.Context, id int) (int, error) {
	todoId, err := tu.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return todoId, nil
}

func (tu *taskUsecase) Update(ctx context.Context, r *http.Request) (*entity.Task, error) {
	rawId := httprouter.ParamsFromContext(r.Context()).ByName("id")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	task, err := tu.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var input struct {
		Title       *string `json:"title"`
		IsCompleted *bool   `json:"is_completed"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.IsCompleted != nil {
		task.IsCompleted = *input.IsCompleted
	}

	task, err = tu.repo.Update(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
