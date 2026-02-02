package mysql

import (
	"context"
	"errors"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type limitRepository struct{}

func NewLimitRepository() repository.LimitRepository {
	return &limitRepository{}
}

func (r *limitRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateLimitPayload) (int64, error) {

	query := `INSERT INTO credit_limits (requested_limit, consumer_id) VALUES (?, ?)`
	args := []any{payload.Requested, payload.ConsumerId}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil

}

func (r *limitRepository) UpdateStatus(ctx context.Context, exec repository.QueryExecutor, payload entity.UpdateLimitPayload) error {

	query := `UPDATE credit_limits SET status = ? WHERE id = ?`
	args := []any{payload.Action, payload.LimitId}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("update failed, limit not found")
	}

	return nil

}

func (r *limitRepository) GetLimitById(ctx context.Context, exec repository.QueryExecutor, id int) (*entity.Limit, error) {

	var l entity.Limit

	query := `SELECT id, requested_limit, status, approved_by, approved_at, consumer_id FROM credit_limits WHERE id = ?`
	err := exec.QueryRowContext(ctx, query, id).Scan(&l.Id, &l.Requested, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
	if err != nil {
		return nil, err
	}

	return &l, nil

}
