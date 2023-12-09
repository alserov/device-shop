package db

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/user-service/internal/db/models"
)

type UserRepo interface {
	InternalActions
	BalanceActions
	AuthActions
	InfoActions
}

type InfoActions interface {
	GetInfo(context.Context, string) (models.GetUserInfoRes, error)
	GetUserInfo(ctx context.Context, uuid string) (models.GetUserInfoRes, error)
}

type BalanceActions interface {
	TopUpBalance(context.Context, models.BalanceReq) (float32, error)
	DebitBalance(context.Context, models.BalanceReq) (float32, error)
	DebitBalanceTx(context.Context, models.BalanceReq) (*sql.Tx, error)
}

type AuthActions interface {
	Signup(ctx context.Context, req models.SignupReq, info models.SignupInfo) error
	Login(ctx context.Context, req models.LoginReq, rToken string) (string, error)
}

type InternalActions interface {
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
	CheckIfAdmin(ctx context.Context, uuid string) (bool, error)
}
