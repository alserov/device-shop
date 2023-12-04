package db

import (
	"context"
	"github.com/alserov/device-shop/user-service/internal/db/models"
)

type UserRepo interface {
	GetInfo(context.Context, string) (models.GetUserInfoRes, error)
	TopUpBalance(context.Context, models.BalanceReq) (float32, error)
	DebitBalance(context.Context, models.BalanceReq) (float32, error)
	GetUserInfo(ctx context.Context, uuid string) (models.GetUserInfoRes, error)

	// auth
	Signup(ctx context.Context, req models.SignupReq, info models.SignupInfo) error
	Login(ctx context.Context, req models.LoginReq, rToken string) (string, error)

	// internal
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
	CheckIfAdmin(ctx context.Context, uuid string) (bool, error)
}
