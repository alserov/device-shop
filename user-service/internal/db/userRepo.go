package db

import (
	"context"
	pb "github.com/alserov/device-shop/proto/gen"
)

type UserRepo interface {
	GetInfo(context.Context, string) (*pb.GetUserInfoRes, error)
	TopUpBalance(context.Context, *pb.BalanceReq) (float32, error)
	DebitBalance(context.Context, *pb.BalanceReq) (float32, error)
}
