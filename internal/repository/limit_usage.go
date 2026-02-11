package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type LimitUsageRepository interface {
	Create(context.Context, QueryExecutor, entity.CreateLimitUsagePayload) error
}
