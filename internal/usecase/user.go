package usecase

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(context.Context, *entity.RegisterPayload) (*entity.User, error)
	Login(context.Context, *entity.LoginPayload) (string, error)
}

type userUsecase struct {
	db *sql.DB
	ur repository.UserRepository
	cr repository.ConsumerRepository
}

func NewUserUsecase(db *sql.DB, ur repository.UserRepository, cr repository.ConsumerRepository) *userUsecase {
	return &userUsecase{
		db: db,
		ur: ur,
		cr: cr,
	}
}

func (uc *userUsecase) Register(ctx context.Context, payload *entity.RegisterPayload) (*entity.User, error) {
	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		return nil, err
	}
	payload.Password = string(hashed)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return uc.ur.Create(ctx, uc.db, *payload)
}

func (uc *userUsecase) Login(ctx context.Context, payload *entity.LoginPayload) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := uc.ur.FindByEmail(ctx, uc.db, payload.Email)
	if err != nil {
		return "", err
	}
	var consumerId int
	if user.RoleId == utils.Roles["admin"] {
		consumerId = 0
	} else {
		consumerId, err = uc.cr.GetIdByUserId(ctx, uc.db, user.Id)
		log.Printf("CONSUMER_ID %d", consumerId)
		if err != nil {
			return "", err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", err
	}

	usr := utils.JwtPayload{
		Id:         user.Id,
		Email:      user.Email,
		RoleId:     user.RoleId,
		ConsumerId: consumerId,
	}

	token := utils.GenerateToken(usr)

	return token, nil

}
