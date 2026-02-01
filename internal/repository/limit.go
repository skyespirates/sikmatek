package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type LimitRepository interface {
	Create(context.Context, entity.CreateLimitPayload) (*entity.Limit, error)
	Action(context.Context, entity.UpdateLimitPayload) (*entity.Limit, error)
}
