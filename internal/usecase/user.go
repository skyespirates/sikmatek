package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/skyespirates/sikmatek/internal/entity"
	"github.com/skyespirates/sikmatek/internal/repository"
	"github.com/skyespirates/sikmatek/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

var ErrNotFound = errors.New("incorrect username or password")

type UserUsecase interface {
	Register(context.Context, *entity.RegisterPayload) (string, error)
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

func (uc *userUsecase) Register(ctx context.Context, payload *entity.RegisterPayload) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := uc.db.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		return "", err
	}
	payload.Password = string(hashed)

	user, err := uc.ur.Create(ctx, uc.db, *payload)
	if err != nil {
		return "", err
	}

	consumerId, err := uc.cr.Create(ctx, tx, user.Id)
	if err != nil {
		return "", err
	}

	isVerified, err := uc.cr.GetIsVerifiedById(ctx, tx, consumerId)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	jwtPayload := utils.JwtPayload{
		Id:         user.Id,
		Email:      user.Email,
		RoleId:     user.RoleId,
		ConsumerId: consumerId,
		IsVerified: isVerified,
	}

	token := utils.GenerateToken(jwtPayload)

	return token, nil

}

func (uc *userUsecase) Login(ctx context.Context, payload *entity.LoginPayload) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := uc.ur.FindByEmail(ctx, uc.db, payload.Email)
	if err != nil {
		return "", err
	}

	var consumerId int
	var isVerified bool
	if user.RoleId == utils.Roles["admin"] {
		consumerId = 0
		isVerified = false
	} else {
		consumer, err := uc.cr.GetByUserId(ctx, uc.db, user.Id)
		if err != nil {
			return "", err
		}
		consumerId = consumer.Id
		isVerified, err = uc.cr.GetIsVerifiedById(ctx, uc.db, consumer.Id)
		if err != nil {
			return "", err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", ErrNotFound
	}

	usr := utils.JwtPayload{
		Id:         user.Id,
		Email:      user.Email,
		RoleId:     user.RoleId,
		ConsumerId: consumerId,
		IsVerified: isVerified,
	}

	token := utils.GenerateToken(usr)

	return token, nil

}
