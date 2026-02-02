package usecase

import (
	"context"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
)

type ContractUsecase interface {
	Create(context.Context, entity.CreateContractPayload) (string, error)
	Quote(context.Context) error
	Confirm(context.Context) error
	Cancel(context.Context) error
	Activate(context.Context) error
	Cicil(context.Context) error
	Detail(context.Context) error
	DaftarKontrak(context.Context) error
}

type contractUsecase struct {
	repo repository.ContractRepository
}

func NewContractUsecase(repo repository.ContractRepository) ContractUsecase {
	return &contractUsecase{
		repo: repo,
	}
}

func (uc *contractUsecase) Create(ctx context.Context, payload entity.CreateContractPayload) (string, error) {
	nomor_kontrak := "INI-NOMOR-KONTRAK-YAA"

	return nomor_kontrak, nil
}

func (uc *contractUsecase) Quote(ctx context.Context) error {
	return nil
}

func (uc *contractUsecase) Confirm(ctx context.Context) error {
	return nil
}

func (uc *contractUsecase) Cancel(ctx context.Context) error {
	return nil
}

func (uc *contractUsecase) Activate(ctx context.Context) error {
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
