package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type TaskRepository interface {
	GetAll(context.Context) ([]*entity.Task, error)
	GetById(context.Context, int) (*entity.Task, error)
	Create(context.Context, string) (*entity.Task, error)
	Delete(context.Context, int) (int, error)
	Update(context.Context, *entity.Task) (*entity.Task, error)
}
