package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type InstallmentRepository interface {
	CreateN(context.Context, QueryExecutor, entity.CreateInstallmentPayload) error
}
