package mysql

import (
	"context"
	"errors"
	"log"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
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

	log.Println("*****************", id)

	var l entity.Limit

	query := `SELECT id, requested_limit, status, approved_by, approved_at, consumer_id FROM credit_limits WHERE id = ?`
	err := exec.QueryRowContext(ctx, query, id).Scan(&l.Id, &l.Requested, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
	if err != nil {
		return nil, err
	}

	return &l, nil

}

func (r *limitRepository) GetLimitList(ctx context.Context, exec repository.QueryExecutor, payload entity.LimitListPayload) ([]*entity.LimitDetail, error) {

	// role["admin"] = 1
	// role["consumer"] = 2

	query := `
		SELECT 
			cl.id, 
			cl.requested_limit, 
			COALESCE(SUM(lu.used_amount), 0) as used_limit, 
			CASE 
				WHEN cl.status = 'APPROVED'
					THEN requested_limit - COALESCE(SUM(lu.used_amount), 0)
				ELSE 0
			END AS remaining_limit, 
			cl.status, 
			cl.approved_by, 
			cl.approved_at, 
			cl.consumer_id 
		FROM credit_limits cl 
		LEFT JOIN limit_usages lu ON
			cl.id = lu.limit_id
	`

	args := []any{}

	if payload.RoleId == utils.Roles["consumer"] {
		query += `WHERE cl.consumer_id = ? `
		args = append(args, payload.ConsumerId)
	}

	query += `GROUP BY cl.id, cl.requested_limit, cl.status, cl.approved_by, cl.approved_at, cl.consumer_id`

	rows, err := exec.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var limits []*entity.LimitDetail
	for rows.Next() {

		var l entity.LimitDetail

		err = rows.Scan(&l.Id, &l.Requested, &l.Used, &l.Remaining, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
		if err != nil {
			return nil, err
		}

		limits = append(limits, &l)

	}

	return limits, nil

}

func (r *limitRepository) GetActiveLimit(ctx context.Context, exec repository.QueryExecutor, consumer_id int) ([]*entity.Limit, error) {

	query := `SELECT  id, requested_limit, status, approved_by, approved_at, consumer_id FROM credit_limits WHERE status = 'APPROVED' AND consumer_id = ?`

	rows, err := exec.QueryContext(ctx, query, consumer_id)
	if err != nil {
		return nil, err
	}

	var limits []*entity.Limit

	for rows.Next() {
		var l entity.Limit

		err = rows.Scan(&l.Id, &l.Requested, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
		if err != nil {
			return nil, err
		}

		limits = append(limits, &l)
	}

	return limits, nil

}
