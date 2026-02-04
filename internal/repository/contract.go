package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type ContractRepository interface {
	Create(context.Context, QueryExecutor, entity.CreateContractPayload) (string, error)
	Update(context.Context, QueryExecutor, entity.UpdateContractPayload) error
	List(context.Context, QueryExecutor, entity.ListContractPayload) ([]*entity.Contract, error)
	Quote(context.Context, QueryExecutor, entity.QuoteContractPayload) error
	GetByNomorKontrak(context.Context, QueryExecutor, string) (*entity.Contract, error)
	ConsumerAction(context.Context, QueryExecutor, entity.ConsumerActionPayload) error
}
