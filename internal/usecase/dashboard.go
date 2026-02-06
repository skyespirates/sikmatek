package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
	"golang.org/x/sync/errgroup"
)

type DashboardUsecase interface {
	GetConsumerDashboardData(context.Context) (map[string]any, error)
}

type dashboardUsecase struct {
	db *sql.DB
	cr repository.ConsumerRepository
	lr repository.LimitRepository
	kr repository.ContractRepository
	pr repository.ProductRepository
}

func NewDashboardUsecase(db *sql.DB, cr repository.ConsumerRepository, lr repository.LimitRepository, kr repository.ContractRepository, pr repository.ProductRepository) DashboardUsecase {
	return &dashboardUsecase{
		db: db,
		cr: cr,
		lr: lr,
		kr: kr,
		pr: pr,
	}
}

func (uc *dashboardUsecase) GetConsumerDashboardData(ctx context.Context) (map[string]any, error) {

	data := map[string]any{}

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	// prepare payload
	claims := utils.ContextGetUser(ctx)

	limitPayload := entity.LimitListPayload{
		RoleId:     claims.RoleId,
		ConsumerId: claims.ConsumerId,
	}

	contractPayload := entity.ListContractPayload{
		RoleId:     claims.RoleId,
		ConsumerId: claims.ConsumerId,
	}

	var (
		consumer  *entity.Consumer
		limits    []*entity.Limit
		contracts []*entity.Contract
		products  []*entity.Product
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		c, err := uc.cr.GetByUserId(ctx, uc.db, claims.Id)
		if err != nil {
			return err
		}

		consumer = c
		return nil
	})

	g.Go(func() error {

		l, err := uc.lr.GetLimitList(ctx, uc.db, limitPayload)
		if err != nil {
			return err
		}

		limits = l
		return nil

	})

	g.Go(func() error {
		c, err := uc.kr.List(ctx, uc.db, contractPayload)
		if err != nil {
			return err
		}

		contracts = c
		return nil

	})

	g.Go(func() error {
		p, err := uc.pr.GetProductList(ctx, uc.db)
		if err != nil {
			return err
		}

		products = p
		return nil

	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	data["profile_info"] = consumer
	data["limits"] = limits
	data["contracts"] = contracts
	data["products"] = products

	return data, nil

}
