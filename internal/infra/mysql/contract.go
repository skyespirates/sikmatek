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
