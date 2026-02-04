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

	var status_limit string
	query := `SELECT status FROM credit_limits WHERE id = ?`
	err := exec.QueryRowContext(ctx, query, payload.LimitId).Scan(&status_limit)
	if err != nil {
		return "", err
	}

	if status_limit != "APPROVED" {
		return "", errors.New("please provide approved limit")
	}

	var product_info struct {
		nama_produk string
		harga       int
		kategori    string
	}
	query = `SELECT nama, harga, kategory FROM products WHERE id = ?`
	err = exec.QueryRowContext(ctx, query, payload.ProductId).Scan(&product_info.nama_produk, &product_info.harga, &product_info.kategori)
	if err != nil {
		return "", err
	}

	nomor_kontrak := utils.GenerateContractID(product_info.kategori)

	query = `INSERT INTO contracts (nomor_kontrak, otr, tenor, consumer_id, product_id, limit_id) VALUES (?, ?, ?, ?, ?, ?)`
	args := []any{nomor_kontrak, product_info.harga, payload.Tenor, payload.ConsumerId, payload.ProductId, payload.LimitId}
	_, err = exec.ExecContext(ctx, query, args...)
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
