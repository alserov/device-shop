package db

import (
	"context"
)

type UserRepo interface {
	GetInfo(context.Context, string) (GetUserInfoRes, error)
	TopUpBalance(context.Context, BalanceReq) (float32, error)
	DebitBalance(context.Context, BalanceReq) (float32, error)
}

type GetUserInfoRes struct {
	Username string
	Email    string
	UUID     string
	Cash     float32
}

type BalanceReq struct {
	Cash     float32
	UserUUID string
}
