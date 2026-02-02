package mysql

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type productRepository struct{}

func NewProductRepository() repository.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateProductPayload) error {
	return nil
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
	return nil, nil
}
