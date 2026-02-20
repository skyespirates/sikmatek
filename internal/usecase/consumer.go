package usecase

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type ConsumerUsecase interface {
	GetInfo(context.Context) (*entity.Consumer, error)
	CompleteInfo(context.Context, entity.UpdateConsumerPayload) error
	SetKtp(context.Context, int, string) error
	SetSelfie(context.Context, int, string) error
	Verify(context.Context, string) error
}

type consumerUsecase struct {
	db *sql.DB
	cr repository.ConsumerRepository
}

func NewConsumerUsecase(db *sql.DB, cr repository.ConsumerRepository) ConsumerUsecase {
	return &consumerUsecase{
		db: db,
		cr: cr,
	}
}

func (uc *consumerUsecase) GetInfo(ctx context.Context) (*entity.Consumer, error) {

	claims := utils.ContextGetUser(ctx)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	consumer, err := uc.cr.GetByUserId(ctx, uc.db, claims.Id)
	if err != nil {
		return nil, err
	}

	return consumer, nil

}

func (uc *consumerUsecase) CompleteInfo(ctx context.Context, payload entity.UpdateConsumerPayload) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	log.Println(payload.TanggalLahir)

	// retrieve consumer info from context
	claims := utils.ContextGetUser(ctx)

	consumserId := claims.ConsumerId
	// perform update consumer
	err := uc.cr.Update(ctx, uc.db, consumserId, payload)
	if err != nil {
		return err
	}

	return nil

}

func (uc *consumerUsecase) SetKtp(ctx context.Context, consumerID int, path string) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return uc.cr.SetKtpPath(ctx, uc.db, consumerID, path)

}

func (uc *consumerUsecase) SetSelfie(ctx context.Context, consumerID int, path string) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return uc.cr.SetSelfiePath(ctx, uc.db, consumerID, path)

}

func (uc *consumerUsecase) Verify(ctx context.Context, consumerId string) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	consumerID, err := strconv.Atoi(consumerId)
	if err != nil {
		return err
	}

	return uc.cr.Verify(ctx, uc.db, consumerID)

}
