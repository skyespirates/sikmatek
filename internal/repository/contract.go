package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type ContractRepository interface {
	Create(context.Context, QueryExecutor, entity.CreateContractPayload) (string, error)
	Update(context.Context, QueryExecutor, entity.UpdateContractPayload) error
}
