package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type limitRepository struct {
	db *sql.DB
}

func NewLimitRepository(db *sql.DB) repository.LimitRepository {
	return &limitRepository{
		db: db,
	}
}

func (r *limitRepository) Create(ctx context.Context, payload entity.CreateLimitPayload) (*entity.Limit, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO credit_limits (requested_limit, consumer_id) VALUES (?, ?)`
	args := []any{payload.Requested, payload.ConsumerId}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var l entity.Limit
	query = `SELECT id, requested_limit, status, approved_by, approved_at, consumer_id FROM credit_limits WHERE id = ?`
	err = tx.QueryRowContext(ctx, query, id).Scan(&l.Id, &l.Requested, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *limitRepository) Action(ctx context.Context, payload entity.UpdateLimitPayload) (*entity.Limit, error) {
	log.Printf("payload: %+v", payload)
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `UPDATE credit_limits SET status = ? WHERE id = ?`
	args := []any{payload.Action, payload.LimitId}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("update failed, limit not found")
	}

	var l entity.Limit
	query = `SELECT id, requested_limit, status, approved_by, approved_at, consumer_id FROM credit_limits WHERE id = ?`
	err = tx.QueryRowContext(ctx, query, payload.LimitId).Scan(&l.Id, &l.Requested, &l.Status, &l.ApprovedBy, &l.ApprovedAt, &l.ConsumerId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &l, nil
}
