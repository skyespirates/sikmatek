package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type ContractUsecase interface {
	Create(context.Context, entity.CreateContractPayload) (string, error)
	GenerateQuote(context.Context, string) error
	Confirm(context.Context, string) error
	Cancel(context.Context, string) error
	Activate(context.Context, string) error
	Cicil(context.Context) error
	Detail(context.Context) error
	DaftarKontrak(context.Context) error
}

type contractUsecase struct {
	db *sql.DB
	cr repository.ContractRepository
	lr repository.LimitRepository
	pr repository.ProductRepository
	ir repository.InstallmentRepository
}

func NewContractUsecase(db *sql.DB, cr repository.ContractRepository, lr repository.LimitRepository, pr repository.ProductRepository, ir repository.InstallmentRepository) ContractUsecase {
	return &contractUsecase{
		db: db,
		cr: cr,
		lr: lr,
		pr: pr,
		ir: ir,
	}
}

func (uc *contractUsecase) Create(ctx context.Context, payload entity.CreateContractPayload) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	claims := utils.ContextGetUser(ctx)
	payload.ConsumerId = claims.ConsumerId

	tx, err := uc.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// check apakah limit yg diinput sudah diapprove
	limit, err := uc.lr.GetLimitById(ctx, tx, payload.LimitId)
	if err != nil {
		return "", err
	}

	if limit.Status != "APPROVED" {
		return "", errors.New("please provide approved limit")
	}

	product, err := uc.pr.GetProductById(ctx, tx, payload.ProductId)
	if err != nil {
		return "", err
	}

	if limit.Requested < product.Harga {
		return "", errors.New("insufficient limit")
	}

	payload.Otr = product.Harga
	payload.ProductCategory = product.Kategori

	// buat kontrak
	nomor_kontrak, err := uc.cr.Create(ctx, tx, payload)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return nomor_kontrak, nil

}

func (uc *contractUsecase) GenerateQuote(ctx context.Context, nomor_kontrak string) error {

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

	var payload entity.QuoteContractPayload

	contract, err := uc.cr.GetByNomorKontrak(ctx, tx, nomor_kontrak)
	if err != nil {
		return err
	}

	payload.NomorKontrak = contract.NomorKontrak

	adminFee := int(float64(contract.Otr) * 0.05)
	bunga := int(float64(contract.Otr) * 0.02 * float64(contract.Tenor))

	payload.AdminFee = adminFee
	payload.JumlahBunga = bunga

	err = uc.cr.Quote(ctx, tx, payload)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (uc *contractUsecase) Confirm(ctx context.Context, nomor_kontrak string) error {

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

	// to confirm, current status should be QUOTED
	if contract.Status != "QUOTED" {
		return errors.New("contract must be quoted before confirmation")
	}

	var payload entity.ConsumerActionPayload

	payload.NomorKontrak = nomor_kontrak
	payload.Action = "CONFIRMED"

	err = uc.cr.ConsumerAction(ctx, tx, payload)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (uc *contractUsecase) Cancel(ctx context.Context, nomor_kontrak string) error {

	// quoted atau belum quoted, consumer boleh cancel kok, sans
	// untuk ngecancel, status mah ngga penting
	var payload entity.ConsumerActionPayload

	payload.NomorKontrak = nomor_kontrak
	payload.Action = "CANCELLED"

	return uc.cr.ConsumerAction(ctx, uc.db, payload)

}

func (uc *contractUsecase) Activate(ctx context.Context, nomor_kontrak string) error {

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

	// to confirm, current status should be QUOTED
	if contract.Status != "CONFIRMED" {
		return errors.New("contract must be confirmed before activation")
	}

	var payload entity.ConsumerActionPayload

	payload.NomorKontrak = nomor_kontrak
	payload.Action = "ACTIVE"

	err = uc.cr.ConsumerAction(ctx, tx, payload)
	if err != nil {
		return err
	}

	create_installment_payload := entity.CreateInstallmentPayload{
		NomorKontrak:    contract.NomorKontrak,
		TotalPembiayaan: *contract.TotalPembiayaan,
		Tenor:           contract.Tenor,
		StartDate:       time.Now(),
	}

	err = uc.ir.CreateN(ctx, tx, create_installment_payload)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (uc *contractUsecase) Cicil(ctx context.Context) error {
	return nil
}

func (uc *contractUsecase) Detail(ctx context.Context) error {
	return nil
}

func (uc *contractUsecase) DaftarKontrak(ctx context.Context) error {
	return nil
}
