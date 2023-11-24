package db

import (
	"context"
	"github.com/alserov/device-shop/auth-service/internal/entity"
	pb "github.com/alserov/device-shop/proto/gen"
)

type AuthRepo interface {
	Signup(ctx context.Context, req *pb.SignupReq, info *entity.SignupAdditional) error
	Login(ctx context.Context, req *pb.LoginReq, rToken string) (string, error)
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
}
