package usecase

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type TenorUsecase interface {
	Create(context.Context, int) (*int64, error)
	GetAll(context.Context) ([]*entity.Tenor, error)
}

type tenorUsecase struct {
	repo repository.TenorRepository
}

func NewTenorUsecase(repo repository.TenorRepository) TenorUsecase {
	return &tenorUsecase{
		repo: repo,
	}
}

func (uc *tenorUsecase) Create(ctx context.Context, tenor int) (*int64, error) {
	return uc.repo.Create(ctx, tenor)
}

func (uc *tenorUsecase) GetAll(ctx context.Context) ([]*entity.Tenor, error) {
	return uc.repo.GetList(ctx)
}
