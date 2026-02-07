package usecase

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
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

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	claims := utils.ContextGetUser(ctx)
	log.Printf("claims %+v", claims)

	data := map[string]any{}

	var mu sync.Mutex
	var wg sync.WaitGroup

	runTask := func(key string, task func() (any, error)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := task()

			mu.Lock()
			if err != nil {
				log.Println(err.Error())
				data[key+"_error"] = err.Error()
			} else {
				data[key] = res
			}
			mu.Unlock()
		}()
	}

	limitPayload := entity.LimitListPayload{
		RoleId:     claims.RoleId,
		ConsumerId: claims.ConsumerId,
	}

	contractPayload := entity.ListContractPayload{
		RoleId:     claims.RoleId,
		ConsumerId: claims.ConsumerId,
	}

	runTask("profile_info", func() (any, error) {
		return uc.cr.GetByUserId(ctx, uc.db, claims.Id)
	})

	runTask("limits", func() (any, error) {
		return uc.lr.GetLimitList(ctx, uc.db, limitPayload)
	})

	runTask("contracts", func() (any, error) {
		return uc.kr.List(ctx, uc.db, contractPayload)
	})

	runTask("products", func() (any, error) {
		return uc.pr.GetProductList(ctx, uc.db)
	})

	wg.Wait()

	return data, nil

}
