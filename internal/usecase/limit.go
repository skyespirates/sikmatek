package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type LimitUsecase interface {
	AjukanLimit(context.Context, entity.CreateLimitPayload) (int64, error)
	TindakLanjut(context.Context, entity.UpdateLimitPayload) error
}

type limitUsecase struct {
	db   *sql.DB
	repo repository.LimitRepository
}

func NewLimitUsecase(db *sql.DB, repo repository.LimitRepository) LimitUsecase {
	return &limitUsecase{
		db:   db,
		repo: repo,
	}
}

func (uc *limitUsecase) AjukanLimit(ctx context.Context, payload entity.CreateLimitPayload) (int64, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := uc.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	id, err := uc.repo.Create(ctx, tx, payload)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (uc *limitUsecase) TindakLanjut(ctx context.Context, payload entity.UpdateLimitPayload) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := uc.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = uc.repo.UpdateStatus(ctx, tx, payload)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
