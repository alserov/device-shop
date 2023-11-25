package db

import (
	"context"
	"time"
)

type AuthRepo interface {
	Signup(ctx context.Context, req SignupReq, info SignupInfo) error
	Login(ctx context.Context, req LoginReq, rToken string) (string, error)
	GetPasswordAndRoleByUsername(ctx context.Context, uname string) (string, string, error)
}

type SignupReq struct {
	Username string
	Password string
	Email    string
}

type LoginReq struct {
	Username string
	Password string
}

type SignupInfo struct {
	UUID         string
	Cash         float32
	RefreshToken string
	Role         string
	CreatedAt    *time.Time
}
