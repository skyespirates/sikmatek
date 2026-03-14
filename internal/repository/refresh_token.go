package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type RefreshTokenRepository interface {
	Insert(context.Context, int, string) error
	Get(context.Context, string) (*entity.RefreshToken, error)
}
