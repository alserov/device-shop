package converter

import (
	"github.com/alserov/device-shop/proto/gen/user"
	"github.com/alserov/device-shop/user-service/internal/service/models"
)

type serverAuth struct{}
type ServerAuth interface {
	SignupReqToService(req *user.SignupReq) models.SignupReq
	SignupResToPb(res models.SignupRes) *user.SignupRes
	LoginReqToService(req *user.LoginReq) models.LoginReq
	LoginResToPb(res models.LoginRes) *user.LoginRes
}

type serverInfo struct{}
type ServerInfo interface {
	GetUserInfoResToPb(res models.GetUserInfoRes) *user.GetUserInfoRes
	CheckIfAdminResToPb(isAdmin bool) *user.CheckIfAdminRes
}

type serverBalance struct{}
type ServerBalance interface {
	BalanceReqToService(req *user.BalanceReq) models.BalanceReq
	BalanceResToPb(balance float32) *user.BalanceRes
}

func NewServerConverter() *ServerConverter {
	return &ServerConverter{
		Auth:    &serverAuth{},
		Info:    &serverInfo{},
		Balance: &serverBalance{},
	}
}

type ServerConverter struct {
	Auth    ServerAuth
	Info    ServerInfo
	Balance ServerBalance
}

func (*serverAuth) SignupReqToService(req *user.SignupReq) models.SignupReq {
	return models.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (*serverAuth) SignupResToPb(res models.SignupRes) *user.SignupRes {
	return &user.SignupRes{
		Username:     res.Username,
		UUID:         res.UUID,
		Cash:         res.Cash,
		Email:        res.Email,
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	}
}

func (*serverAuth) LoginReqToService(req *user.LoginReq) models.LoginReq {
	return models.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func (*serverAuth) LoginResToPb(res models.LoginRes) *user.LoginRes {
	return &user.LoginRes{
		RefreshToken: res.RefreshToken,
		Token:        res.Token,
		UUID:         res.UUID,
	}
}

func (*serverInfo) GetUserInfoResToPb(res models.GetUserInfoRes) *user.GetUserInfoRes {
	return &user.GetUserInfoRes{
		Username: res.Username,
		Email:    res.Email,
		UUID:     res.UUID,
		Cash:     res.Cash,
	}
}

func (*serverInfo) CheckIfAdminResToPb(isAdmin bool) *user.CheckIfAdminRes {
	return &user.CheckIfAdminRes{
		IsAdmin: isAdmin,
	}
}

func (*serverBalance) BalanceReqToService(req *user.BalanceReq) models.BalanceReq {
	return models.BalanceReq{
		Cash:     req.Cash,
		UserUUID: req.UserUUID,
	}
}

func (*serverBalance) BalanceResToPb(balance float32) *user.BalanceRes {
	return &user.BalanceRes{
		Cash: balance,
	}
}
