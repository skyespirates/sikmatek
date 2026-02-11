package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type InstallmentUsecase interface {
	GenerateInstallment(context.Context, string) error
	PayInstallment(context.Context, int) error
}

type installmentUsecase struct {
	db *sql.DB
	ir repository.InstallmentRepository
	cr repository.ContractRepository
	lr repository.LimitUsageRepository
}

func NewInstallmentUsecase(db *sql.DB, ir repository.InstallmentRepository, cr repository.ContractRepository, lr repository.LimitUsageRepository) InstallmentUsecase {
	return &installmentUsecase{
		db: db,
		ir: ir,
		cr: cr,
		lr: lr,
	}
}

func (uc *installmentUsecase) GenerateInstallment(ctx context.Context, nomor_kontrak string) error {

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

	contract, err := uc.cr.GetByNomorKontrak(ctx, tx, nomor_kontrak)
	if err != nil {
		return err
	}

	payload := entity.CreateInstallmentPayload{
		NomorKontrak:    contract.NomorKontrak,
		TotalPembiayaan: *contract.TotalPembiayaan,
		Tenor:           contract.Tenor,
	}
	err = uc.ir.CreateN(ctx, tx, payload)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (uc *installmentUsecase) PayInstallment(ctx context.Context, id int) error {

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
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

	// get jumlah tagihan & limit_id
	installment, err := uc.ir.GetInfo(ctx, tx, id)
	if err != nil {
		return err
	}

	// kalo sudah dibayar, return
	if installment.Status == "PAID" {
		return entity.ErrDuplicatePayment
	}

	// update installment status & paid at
	err = uc.ir.Pay(ctx, tx, id)
	if err != nil {
		return err
	}

	// insert into limit usage
	payload := entity.CreateLimitUsagePayload{
		UsedAmount:    installment.JumlahTagihan,
		InstallmentId: id,
		LimitId:       installment.LimitId,
	}
	err = uc.lr.Create(ctx, tx, payload)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil

}
