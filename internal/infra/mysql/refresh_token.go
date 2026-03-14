package mysql

import (
	"context"
	"database/sql"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type refreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) repository.RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
	}
}

func (r *refreshTokenRepository) Insert(ctx context.Context, user_id int, hashedToken string) error {
	query := `INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 7 DAY))`

	args := []any{user_id, hashedToken}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *refreshTokenRepository) Get(ctx context.Context, hashed string) (*entity.RefreshToken, error) {
	query := `SELECT * FROM refresh_tokens WHERE token_hash = ?`
	var dest entity.RefreshToken
	err := r.db.QueryRowContext(ctx, query, hashed).Scan(&dest.ID, &dest.UserID, &dest.TokenHash, &dest.ExpiresAt, &dest.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}
