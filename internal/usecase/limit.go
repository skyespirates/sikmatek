package usecase

import (
	"context"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type LimitUsecase interface {
	AjukanLimit(context.Context, entity.CreateLimitPayload) (*entity.Limit, error)
	TindakLanjut(context.Context, entity.UpdateLimitPayload) (*entity.Limit, error)
}

type limitUsecase struct {
	repo repository.LimitRepository
}

func NewLimitUsecase(repo repository.LimitRepository) LimitUsecase {
	return &limitUsecase{
		repo: repo,
	}
}

func (u *limitUsecase) AjukanLimit(ctx context.Context, payload entity.CreateLimitPayload) (*entity.Limit, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.repo.Create(ctx, payload)
}

func (u *limitUsecase) TindakLanjut(ctx context.Context, payload entity.UpdateLimitPayload) (*entity.Limit, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.repo.Action(ctx, payload)
}
