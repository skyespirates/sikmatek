package mysql

import (
	"context"
	"errors"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
)

type contractRepository struct{}

func NewContractRepository() repository.ContractRepository {
	return &contractRepository{}
}

func (r *contractRepository) Create(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateContractPayload) (string, error) {

	nomor_kontrak := utils.GenerateContractID(payload.ProductCategory)

	query := `INSERT INTO contracts (nomor_kontrak, otr, tenor, consumer_id, product_id, limit_id) VALUES (?, ?, ?, ?, ?, ?)`
	args := []any{nomor_kontrak, payload.Otr, payload.Tenor, payload.ConsumerId, payload.ProductId, payload.LimitId}
	_, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return nomor_kontrak, nil

}

func (r *contractRepository) Update(ctx context.Context, exec repository.QueryExecutor, payload entity.UpdateContractPayload) error {
	return nil
}

func (r *contractRepository) List(ctx context.Context, exec repository.QueryExecutor, payload entity.ListContractPayload) ([]*entity.Contract, error) {

	var query string
	args := []any{}
	if payload.RoleId == utils.Roles["admin"] {
		query = `SELECT nomor_kontrak, otr, admin_fee, jumlah_bunga, tenor, total_pembiayaan, status, consumer_id, product_id, limit_id FROM contracts`
	} else {
		query = `SELECT nomor_kontrak, otr, admin_fee, jumlah_bunga, tenor, total_pembiayaan, status, consumer_id, product_id, limit_id FROM contracts WHERE consumer_id = ?`
		args = append(args, payload.ConsumerId)
	}

	rows, err := exec.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []*entity.Contract
	for rows.Next() {

		var c entity.Contract

		err = rows.Scan(&c.NomorKontrak, &c.Otr, &c.AdminFee, &c.JumlahBunga, &c.Tenor, &c.TotalPembiayaan, &c.Status, &c.ConsumerId, &c.ProductId, &c.LimitId)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, &c)

	}

	return contracts, nil

}

func (r *contractRepository) Quote(ctx context.Context, exec repository.QueryExecutor, payload entity.QuoteContractPayload) error {

	query := `UPDATE contracts SET admin_fee = ?, jumlah_bunga = ?, status = 'QUOTED' WHERE nomor_kontrak = ?`
	args := []any{payload.AdminFee, payload.JumlahBunga, payload.NomorKontrak}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to generate quote")
	}

	return nil

}

func (r *contractRepository) GetByNomorKontrak(ctx context.Context, exec repository.QueryExecutor, nomor_kontrak string) (*entity.Contract, error) {

	query := `SELECT nomor_kontrak, otr, admin_fee, jumlah_bunga, tenor, total_pembiayaan, status, consumer_id, product_id, limit_id FROM contracts WHERE nomor_kontrak = ?`
	var contract entity.Contract
	err := exec.QueryRowContext(ctx, query, nomor_kontrak).Scan(&contract.NomorKontrak, &contract.Otr, &contract.AdminFee, &contract.JumlahBunga, &contract.Tenor, &contract.TotalPembiayaan, &contract.Status, &contract.ConsumerId, &contract.ProductId, &contract.LimitId)
	if err != nil {
		return nil, err
	}

	return &contract, nil

}

func (r *contractRepository) ConsumerAction(ctx context.Context, exec repository.QueryExecutor, payload entity.ConsumerActionPayload) error {

	query := `UPDATE contracts SET status = ? WHERE nomor_kontrak = ?`
	args := []any{payload.Action, payload.NomorKontrak}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to perform consumer action")
	}

	return nil

}
