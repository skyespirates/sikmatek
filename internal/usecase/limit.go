package usecase

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type LimitUsecase interface {
	GetList(context.Context) ([]*entity.LimitDetail, error)
	AjukanLimit(context.Context, int) (int64, error)
	TindakLanjut(context.Context, entity.UpdateLimitPayload) error
	ListLimitAktif(context.Context) ([]*entity.ApprovedLimit, error)
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

func (uc *limitUsecase) GetList(ctx context.Context) ([]*entity.LimitDetail, error) {

	claims := utils.ContextGetUser(ctx)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	payload := entity.LimitListPayload{
		RoleId:     claims.RoleId,
		ConsumerId: claims.ConsumerId,
	}

	return uc.repo.GetLimitList(ctx, uc.db, payload)

}

func (uc *limitUsecase) AjukanLimit(ctx context.Context, requested_limit int) (int64, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	claims := utils.ContextGetUser(ctx)

	payload := entity.CreateLimitPayload{
		Requested:  requested_limit,
		ConsumerId: claims.ConsumerId,
	}

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

func (uc *limitUsecase) ListLimitAktif(ctx context.Context) ([]*entity.ApprovedLimit, error) {

	claims := utils.ContextGetUser(ctx)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	limits, err := uc.repo.GetActiveLimit(ctx, uc.db, claims.ConsumerId)
	if err != nil {
		return nil, err
	}

	p := message.NewPrinter(language.Indonesian)

	var l []*entity.ApprovedLimit
	for _, v := range limits {
		var temp entity.ApprovedLimit
		temp.Value = strconv.Itoa(v.Id)
		temp.Label = p.Sprintf("%d", v.Requested)

		l = append(l, &temp)
	}

	return l, nil

}
