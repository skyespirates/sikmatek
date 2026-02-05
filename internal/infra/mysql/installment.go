package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/skyespirates/sikmatek/internal/repository"
)

type installmentRepository struct{}

func NewInstallmentRepository() repository.InstallmentRepository {
	return &installmentRepository{}
}

func (r *installmentRepository) CreateN(ctx context.Context, exec repository.QueryExecutor, nomor_kontrak string, tenor int) error {

	placeholders := make([]string, tenor)
	args := make([]any, tenor*2)

	for i := 0; i < tenor; i++ {
		placeholders[i] = "(?, ?)"
		args[i*2] = nomor_kontrak
		args[i*2+1] = i + 1
	}

	str := strings.Join(placeholders, ",")

	query := fmt.Sprintf("INSERT INTO installments (nomor_kontrak, bulan_ke) VALUES %s", str)

	_, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil

}
