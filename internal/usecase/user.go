package usecase

import (
	"context"
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
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *userUsecase {
	return &userUsecase{
		repo: repo,
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

	return uc.repo.Create(ctx, *payload)
}

func (uc *userUsecase) Login(ctx context.Context, payload *entity.LoginPayload) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := uc.repo.FindByEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", err
	}

	usr := utils.JwtPayload{
		Id:         user.Id,
		Email:      user.Email,
		RoleId:     user.RoleId,
		ConsumerId: user.ConsumerId,
	}

	token := utils.GenerateToken(usr)

	return token, nil
}
