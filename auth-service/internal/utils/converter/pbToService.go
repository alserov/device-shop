package converter

import (
	"github.com/alserov/device-shop/auth-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/auth"
)

func SignupReqToServiceStruct(req *auth.SignupReq) *models.SignupReq {
	return &models.SignupReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func SignupResToPb(res *models.SignupRes) *auth.SignupRes {
	return &auth.SignupRes{
		Username:     res.Username,
		UUID:         res.UUID,
		Cash:         res.Cash,
		Email:        res.Email,
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	}
}

func LoginReqToServiceStruct(req *auth.LoginReq) *models.LoginReq {
	return &models.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func LoginResToPb(res *models.LoginRes) *auth.LoginRes {
	return &auth.LoginRes{
		RefreshToken: res.RefreshToken,
		Token:        res.Token,
		UUID:         res.UUID,
	}
}

func GetUserInfoResToPb(res *models.GetUserInfoRes) *auth.GetUserInfoRes {
	return &auth.GetUserInfoRes{
		Username: res.Username,
		Email:    res.Email,
		UUID:     res.UUID,
		Cash:     res.Cash,
	}
}

func CheckIfAdminResToPb(isAdmin bool) *auth.CheckIfAdminRes {
	return &auth.CheckIfAdminRes{
		IsAdmin: isAdmin,
	}
}
