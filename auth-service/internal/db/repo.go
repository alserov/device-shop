package db

import (
	"context"
	"github.com/alserov/device-shop/auth-service/internal/entity"
	"time"
)

type AuthRepo interface {
	Signup(ctx context.Context, req *entity.SignupReq, info SignupInfo) error
	Login(ctx context.Context, req *entity.LoginReq, rToken string) (string, error)
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
	GetUserInfo(context.Context, *entity.GetUserInfoReq) (*entity.GetUserInfoRes, error)
	CheckIfAdmin(context.Context, *entity.CheckIfAdminReq) (*entity.CheckIfAdminRes, error)
}

type SignupInfo struct {
	UUID         string
	Cash         float32
	RefreshToken string
	Role         string
	CreatedAt    *time.Time
}
