package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type TenorRepository interface {
	Create(context.Context, int) (*int64, error)
	GetList(context.Context) ([]*entity.Tenor, error)
}
