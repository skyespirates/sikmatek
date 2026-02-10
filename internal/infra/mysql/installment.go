package mysql

import (
	"context"
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
