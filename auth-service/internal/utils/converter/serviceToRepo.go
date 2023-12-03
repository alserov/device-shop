package converter

import (
	repo "github.com/alserov/device-shop/auth-service/internal/db/models"
	service "github.com/alserov/device-shop/auth-service/internal/service/models"
)

func ServiceSignupReqToRepo(req *service.SignupReq) *repo.SignupReq {
	return &repo.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ServiceLoginReqToRepo(req *service.LoginReq) *repo.LoginReq {
	return &repo.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func RepoUserInfoResToService(res *repo.GetUserInfoRes) *service.GetUserInfoRes {
	return &service.GetUserInfoRes{
		UUID:     res.UUID,
		Username: res.Username,
		Email:    res.Email,
		Cash:     res.Cash,
	}
}
