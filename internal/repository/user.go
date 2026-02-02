package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type UserRepository interface {
	Create(context.Context, QueryExecutor, entity.RegisterPayload) (*entity.User, error)
	FindByEmail(context.Context, QueryExecutor, string) (*entity.User, error)
}
