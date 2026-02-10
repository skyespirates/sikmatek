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
}

type installmentUsecase struct {
	db *sql.DB
	ir repository.InstallmentRepository
	cr repository.ContractRepository
}

func NewInstallmentUsecase(db *sql.DB, ir repository.InstallmentRepository, cr repository.ContractRepository) InstallmentUsecase {
	return &installmentUsecase{
		db: db,
		ir: ir,
		cr: cr,
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
