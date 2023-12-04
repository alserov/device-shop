package converter

import (
	repo "github.com/alserov/device-shop/user-service/internal/db/models"
	"github.com/alserov/device-shop/user-service/internal/service/models"
)

type ServiceConverter struct {
	Auth ServiceAuth
	Info ServiceInfo
}

func NewServiceConverter() *ServiceConverter {
	return &ServiceConverter{
		Auth: &serviceAuth{},
		Info: &serviceInfo{},
	}
}

type serviceAuth struct{}

type ServiceAuth interface {
	SignupReqToRepo(req models.SignupReq) repo.SignupReq
	SignupResToService(res models.SignupReq, info repo.SignupInfo, token string) models.SignupRes
	LoginReqToRepo(req models.LoginReq) repo.LoginReq
	LoginResToService(rToken string, token string, uuid string) models.LoginRes
}

type serviceInfo struct{}
type ServiceInfo interface {
	UserInfoResToService(res repo.GetUserInfoRes) models.GetUserInfoRes
}

func (*serviceAuth) SignupReqToRepo(req models.SignupReq) repo.SignupReq {
	return repo.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (a *serviceAuth) SignupResToService(res models.SignupReq, info repo.SignupInfo, token string) models.SignupRes {
	return models.SignupRes{
		Username:     res.Username,
		Email:        res.Email,
		UUID:         info.UUID,
		Cash:         info.Cash,
		RefreshToken: info.RefreshToken,
		Token:        token,
	}
}

func (*serviceAuth) LoginReqToRepo(req models.LoginReq) repo.LoginReq {
	return repo.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func (*serviceAuth) LoginResToService(rToken string, token string, uuid string) models.LoginRes {
	return models.LoginRes{
		RefreshToken: rToken,
		Token:        token,
		UUID:         uuid,
	}
}

func (*serviceInfo) UserInfoResToService(res repo.GetUserInfoRes) models.GetUserInfoRes {
	return models.GetUserInfoRes{
		UUID:     res.UUID,
		Username: res.Username,
		Email:    res.Email,
		Cash:     res.Cash,
	}
}
