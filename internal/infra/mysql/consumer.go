package mysql

import (
	"context"
	"errors"

	msql "github.com/go-sql-driver/mysql"
	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

var ErrDuplicateNik = errors.New("nik already used")

type consumerRepository struct{}

func NewConsumerRepository() repository.ConsumerRepository {
	return &consumerRepository{}
}

func (r *consumerRepository) List(ctx context.Context, exec repository.QueryExecutor) ([]*entity.Consumer, error) {

	var consumers []*entity.Consumer

	return consumers, nil

}

func (r *consumerRepository) GetById(ctx context.Context, exec repository.QueryExecutor, id int) (*entity.Consumer, error) {

	var consumer *entity.Consumer

	return consumer, nil

}

func (r *consumerRepository) GetByUserId(ctx context.Context, exec repository.QueryExecutor, user_id int) (*entity.Consumer, error) {

	var c entity.Consumer

	query := `SELECT id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, is_verified, user_id FROM consumers WHERE user_id = ?`
	err := exec.QueryRowContext(ctx, query, user_id).Scan(&c.Id, &c.Nik, &c.FullName, &c.LegalName, &c.TempatLahir, &c.TanggalLahir, &c.Gaji, &c.FotoKtp, &c.FotoSelfie, &c.IsVerified, &c.UserId)
	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (r *consumerRepository) Create(ctx context.Context, exec repository.QueryExecutor, user_id int) (int, error) {

	query := `INSERT INTO consumers (user_id) VALUES (?)`
	result, err := exec.ExecContext(ctx, query, user_id)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (r *consumerRepository) Update(ctx context.Context, exec repository.QueryExecutor, consumer_id int, payload entity.UpdateConsumerPayload) error {

	query := `UPDATE consumers SET nik = ?,  full_name = ?, legal_name = ?, tempat_lahir = ?, tanggal_lahir = ?, gaji = ? WHERE id = ?`
	args := []any{payload.Nik, payload.FullName, payload.LegalName, payload.TempatLahir, payload.TanggalLahir, payload.Gaji, consumer_id}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		if mysqlErr, ok := err.(*msql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrDuplicateNik
		}
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to update consumer")
	}

	return nil

}

func (r *consumerRepository) SetKtpPath(ctx context.Context, exec repository.QueryExecutor, consumerID int, path string) error {

	query := `UPDATE consumers SET foto_ktp = ? WHERE id = ?`
	args := []any{path, consumerID}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to set ktp path")
	}

	return nil

}

func (r *consumerRepository) SetSelfiePath(ctx context.Context, exec repository.QueryExecutor, consumerID int, path string) error {

	query := `UPDATE consumers SET foto_selfie = ? WHERE id = ?`
	args := []any{path, consumerID}

	result, err := exec.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to set selfie path")
	}

	return nil

}

func (r *consumerRepository) Verify(ctx context.Context, exec repository.QueryExecutor, consumerID int) error {

	query := `UPDATE consumers SET is_verified = 1 WHERE id = ?`

	result, err := exec.ExecContext(ctx, query, consumerID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("failed to verified consumer")
	}

	return nil

}

func (r *consumerRepository) GetIsVerifiedById(ctx context.Context, exec repository.QueryExecutor, consumerID int) (bool, error) {

	var isVerified int

	query := `SELECT is_verified FROM consumers WHERE id = ?`
	err := exec.QueryRowContext(ctx, query, consumerID).Scan(&isVerified)
	if err != nil {
		return false, err
	}

	return isVerified == 1, nil

}

func (r *consumerRepository) GetIdByUserId(ctx context.Context, exec repository.QueryExecutor, user_id int) (*entity.ConsumerId, error) {

	query := `SELECT id FROM consumers WHERE user_id = ?`

	var cid entity.ConsumerId

	err := exec.QueryRowContext(ctx, query, user_id).Scan(&cid.Id)
	if err != nil {
		return nil, err
	}

	return &cid, nil

}
