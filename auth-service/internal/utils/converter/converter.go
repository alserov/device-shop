package converter

import (
	"github.com/alserov/device-shop/auth-service/internal/entity"
	"github.com/alserov/device-shop/proto/gen/auth"
)

func SignupReqToServiceStruct(req *auth.SignupReq) *entity.SignupReq {
	return &entity.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func SignupResToPb(res *entity.SignupRes) *auth.SignupRes {
	return &auth.SignupRes{
		Username:     res.Username,
		UUID:         res.UUID,
		Cash:         res.Cash,
		Email:        res.Email,
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	}
}

func LoginReqToServiceStruct(req *auth.LoginReq) *entity.LoginReq {
	return &entity.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func LoginResToPb(res *entity.LoginRes) *auth.LoginRes {
	return &auth.LoginRes{
		RefreshToken: res.RefreshToken,
		Token:        res.Token,
		UUID:         res.UUID,
	}
}

func GetUserInfoReqToServiceStruct(req *auth.GetUserInfoReq) *entity.GetUserInfoReq {
	return &entity.GetUserInfoReq{
		UUID: req.UserUUID,
	}
}

func GetUserInfoResToPb(res *entity.GetUserInfoRes) *auth.GetUserInfoRes {
	return &auth.GetUserInfoRes{
		Username: res.Username,
		Email:    res.Email,
		UUID:     res.UUID,
		Cash:     res.Cash,
	}
}

func CheckIfAdminReqToServiceStruct(req *auth.CheckIfAdminReq) *entity.CheckIfAdminReq {
	return &entity.CheckIfAdminReq{
		UUID: req.UserUUID,
	}
}

func CheckIfAdminResToPb(res *entity.CheckIfAdminRes) *auth.CheckIfAdminRes {
	return &auth.CheckIfAdminRes{
		IsAdmin: res.IsAdmin,
	}
}
