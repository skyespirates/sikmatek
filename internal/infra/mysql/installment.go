package mysql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type installmentRepository struct{}

func NewInstallmentRepository() repository.InstallmentRepository {
	return &installmentRepository{}
}

func (r *installmentRepository) CreateN(ctx context.Context, exec repository.QueryExecutor, payload entity.CreateInstallmentPayload) error {

	placeholders := make([]string, payload.Tenor)
	args := make([]any, payload.Tenor*4)

	tagihan_bulanan := int(payload.TotalPembiayaan / payload.Tenor)

	for i := 0; i < payload.Tenor; i++ {
		placeholders[i] = "(?, ?, ?, ?)"
		args[i*4] = payload.NomorKontrak
		args[i*4+1] = i + 1
		args[i*4+2] = tagihan_bulanan
		args[i*4+3] = payload.StartDate.AddDate(0, i+1, 0)
	}

	str := strings.Join(placeholders, ",")

	query := fmt.Sprintf("INSERT INTO installments (nomor_kontrak, bulan_ke, jumlah_tagihan, due_date) VALUES %s", str)

	_, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil

}

func (r *installmentRepository) GetInfo(ctx context.Context, exec repository.QueryExecutor, id int) (*entity.InstallmentInfo, error) {

	i := entity.InstallmentInfo{}

	query := `
		SELECT i.jumlah_tagihan, i.status, c.limit_id FROM installments i
		JOIN contracts c ON
			i.nomor_kontrak = c.nomor_kontrak
		WHERE i.id = ?
	`

	err := exec.QueryRowContext(ctx, query, id).Scan(&i.JumlahTagihan, &i.Status, &i.LimitId)
	if err != nil {
		return nil, err
	}

	return &i, nil

}

func (r *installmentRepository) Pay(ctx context.Context, exec repository.QueryExecutor, id int) error {

	query := `UPDATE installments SET status = 'PAID', paid_at = NOW() WHERE id = ?`

	result, err := exec.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return errors.New("failed to update instalment")
	}

	return nil

}
