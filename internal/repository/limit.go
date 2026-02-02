package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type LimitRepository interface {
	Create(context.Context, QueryExecutor, entity.CreateLimitPayload) (int64, error)
	UpdateStatus(context.Context, QueryExecutor, entity.UpdateLimitPayload) error
	GetLimitById(context.Context, QueryExecutor, int) (*entity.Limit, error)
}
