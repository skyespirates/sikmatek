package pgsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type userRepository struct {
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{}
}

var ErrDuplicate = errors.New("email already registered")

func (ur *userRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.RegisterPayload) (*entity.User, error) {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password`
	args := []any{payload.Email, payload.Password}

	var user entity.User

	err := exec.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, ErrDuplicate
			}
		}

		return nil, err
	}
	return &user, nil

}

var ErrNotFound = errors.New("user not found")

func (ur *userRepository) FindByEmail(ctx context.Context, exec repository.QueryExecutor, email string) (*entity.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`

	var u entity.User

	err := exec.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}
