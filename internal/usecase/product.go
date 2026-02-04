package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type ProductUsecase interface {
	Create(context.Context, entity.CreateProductPayload) (int, error)
	GetList(context.Context) ([]*entity.Product, error)
	GetById(context.Context, int) (*entity.Product, error)
}

type productUsecase struct {
	db   *sql.DB
	repo repository.ProductRepository
}

func NewProductUsecase(db *sql.DB, repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		db:   db,
		repo: repo,
	}
}

func (uc *productUsecase) Create(ctx context.Context, payload entity.CreateProductPayload) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return uc.repo.Create(ctx, uc.db, payload)

}

func (uc *productUsecase) GetList(ctx context.Context) ([]*entity.Product, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return uc.repo.GetProductList(ctx, uc.db)

}

func (uc *productUsecase) GetById(ctx context.Context, id int) (*entity.Product, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	product, err := uc.repo.GetProductById(ctx, uc.db, id)
	if err != nil {
		return nil, err
	}

	return product, nil

}
