package db

import (
	"context"
	"github.com/alserov/device-shop/auth-service/internal/db/models"
)

type AuthRepo interface {
	Signup(ctx context.Context, req *models.SignupReq, info models.SignupInfo) error
	Login(ctx context.Context, req *models.LoginReq, rToken string) (string, error)
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
	GetUserInfo(ctx context.Context, uuid string) (*models.GetUserInfoRes, error)
	CheckIfAdmin(ctx context.Context, uuid string) (bool, error)
}
