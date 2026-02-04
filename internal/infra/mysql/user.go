package mysql

import (
	"context"
	"database/sql"
	"errors"

	msql "github.com/go-sql-driver/mysql"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

var ErrDuplicate = errors.New("email already registered")

type userRepository struct{}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(ctx context.Context, exec repository.QueryExecutor, user entity.RegisterPayload) (*entity.User, error) {

	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	args := []any{user.Email, user.Password}
	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {

		if mysqlErr, ok := err.(*msql.MySQLError); ok && mysqlErr.Number == 1062 {
			return nil, ErrDuplicate
		}

		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var u entity.User
	query = `SELECT id, email, role_id FROM users WHERE id = ?`
	err = exec.QueryRowContext(ctx, query, id).Scan(&u.Id, &u.Email, &u.RoleId)
	if err != nil {
		return nil, err
	}

	return &u, nil

}

func (ur *userRepository) FindByEmail(ctx context.Context, exec repository.QueryExecutor, email string) (*entity.User, error) {

	query := `SELECT id, email, password, role_id FROM users WHERE email = ?`
	var user entity.User
	err := exec.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Email, &user.Password, &user.RoleId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
