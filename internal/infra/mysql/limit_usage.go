package mysql

import (
	"context"
	"errors"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type limitUsageRepository struct{}

func NewLimitUsageRepository() repository.LimitUsageRepository {
	return &limitUsageRepository{}
}

func (r *limitUsageRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateLimitUsagePayload) error {

	query := `INSERT INTO limit_usages (used_amount, installment_id, limit_id) VALUES (?, ?, ?)`
	args := []any{payload.UsedAmount, payload.InstallmentId, payload.LimitId}
	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if insertedId == 0 {
		return errors.New("inserted failed: no ID returned")
	}

	return nil

}
