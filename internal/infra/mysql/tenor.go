package mysql

import (
	"context"
	"database/sql"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type tenorRepository struct {
	db *sql.DB
}

func NewTenorRepository(db *sql.DB) repository.TenorRepository {
	return &tenorRepository{
		db: db,
	}
}

func (r *tenorRepository) Create(ctx context.Context, tenor int) (*int64, error) {
	query := `INSERT INTO tenors (durasi_bulan) VALUES (?)`
	args := []interface{}{tenor}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *tenorRepository) GetList(ctx context.Context) ([]*entity.Tenor, error) {
	query := `SELECT id, durasi_bulan FROM tenors`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tenors := []*entity.Tenor{}
	for rows.Next() {
		var t entity.Tenor
		err := rows.Scan(&t.Id, &t.DurasiBulan)
		if err != nil {
			return nil, err
		}

		tenors = append(tenors, &t)
	}

	return tenors, nil
}
