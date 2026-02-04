package repository

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
)

type ProductRepository interface {
	Create(context.Context, QueryExecutor, entity.CreateProductPayload) (int, error)
	GetProductById(context.Context, QueryExecutor, int) (*entity.Product, error)
	GetProductList(context.Context, QueryExecutor) ([]*entity.Product, error)
}
