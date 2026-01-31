package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

var DuplicateErr = errors.New("email already registered")

func (ur *userRepo) Create(ctx context.Context, payload entity.RegisterPayload) (*entity.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, password, created_at, updated_at, version`
	args := []any{payload.Name, payload.Email, payload.Password}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user entity.User

	err := ur.db.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Version)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, DuplicateErr
			}
		}

		return nil, err
	}
	return &user, nil

}

var ErrNotFound = errors.New("user not found")

func (ur *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at, version FROM users WHERE email = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var u entity.User

	err := ur.db.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}
