package mysql

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/repository"
)

type consumerRepository struct{}

func NewConsumerRepository() repository.ConsumerRepository {
	return &consumerRepository{}
}

func (r *consumerRepository) GetIdByUserId(ctx context.Context, exec repository.QueryExecutor, user_id int) (int, error) {

	var id int

	query := `SELECT id FROM consumers WHERE user_id = ?`
	err := exec.QueryRowContext(ctx, query, user_id).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
