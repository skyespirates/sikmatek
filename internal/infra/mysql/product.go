package mysql

import (
	"context"
	"fmt"

	msql "github.com/go-sql-driver/mysql"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type productRepository struct{}

func NewProductRepository() repository.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateProductPayload) (int, error) {

	query := `INSERT INTO products (nama, kategori, harga) VALUES (?, ?, ?)`
	args := []any{payload.Nama, payload.Kategori, payload.Harga}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {

		if mysqlErr, ok := err.(*msql.MySQLError); ok && mysqlErr.Number == 1265 {
			return 0, fmt.Errorf("invalid kategori value: %s", payload.Kategori)
		}

		return 0, err

	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (r *productRepository) GetProductById(ctx context.Context, exec repository.QueryExecutor, id int) (*entity.Product, error) {

	var p entity.Product

	query := `SELECT id, nama, kategori, harga FROM products WHERE id = ?`
	err := exec.QueryRowContext(ctx, query, id).Scan(&p.Id, &p.Nama, &p.Kategori, &p.Harga)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepository) GetProductList(ctx context.Context, exec repository.QueryExecutor) ([]*entity.Product, error) {

	query := `SELECT id, nama, kategori, harga FROM products`

	rows, err := exec.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	products := []*entity.Product{}
	for rows.Next() {
		var p entity.Product

		err = rows.Scan(&p.Id, &p.Nama, &p.Kategori, &p.Harga)
		if err != nil {
			return nil, err
		}

		products = append(products, &p)

	}

	return products, nil

}
