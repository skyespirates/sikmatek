package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type ConsumerRepository interface {
	List(context.Context, QueryExecutor) ([]*entity.Consumer, error)
	GetById(context.Context, QueryExecutor, int) (*entity.Consumer, error)
	GetIdByUserId(context.Context, QueryExecutor, int) (int, error)
	Create(context.Context, QueryExecutor, int) (int, error)
	Update(context.Context, QueryExecutor, int, entity.UpdateConsumerPayload) error
	SetKtpPath(context.Context, QueryExecutor, int, string) error
	SetSelfiePath(context.Context, QueryExecutor, int, string) error
	Verify(context.Context, QueryExecutor, int) error
}
