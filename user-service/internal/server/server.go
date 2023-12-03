package server

import (
	"context"
	"github.com/alserov/device-shop/proto/gen/auth"
	"github.com/alserov/device-shop/proto/gen/user"
	"github.com/alserov/device-shop/user-service/internal/service"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"
	"github.com/alserov/device-shop/user-service/internal/utils/validation"
	"google.golang.org/grpc"
)

func Register(s *grpc.Server) {
	user.RegisterUsersServer(s, &server{})
}

type server struct {
	user.UnimplementedUsersServer
	user service.Service
}

func (s server) TopUpBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DebitBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) Signup(ctx context.Context, req *auth.SignupReq) (*auth.SignupRes, error) {
	if err := validation.ValidateSignupReq(req); err != nil {
		return nil, err
	}

	res, err := s.user.Signup(ctx, converter.SignupReqToServiceStruct(req))
	if err != nil {
		return nil, err
	}

	return converter.SignupResToPb(res), nil
}

func (s server) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginRes, error) {
	if err := validation.ValidateLoginReq(req); err != nil {
		return nil, err
	}

	res, err := s.auth.Login(ctx, converter.LoginReqToServiceStruct(req))
	if err != nil {
		return nil, err
	}

	return converter.LoginResToPb(res), nil
}

func (s server) GetUserInfo(ctx context.Context, req *auth.GetUserInfoReq) (*auth.GetUserInfoRes, error) {
	if err := validation.ValidateGetUserInfoReq(req); err != nil {
		return nil, err
	}

	res, err := s.auth.GetUserInfo(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return converter.GetUserInfoResToPb(res), nil
}

func (s server) CheckIfAdmin(ctx context.Context, req *auth.CheckIfAdminReq) (*auth.CheckIfAdminRes, error) {
	if err := validation.ValidateCheckIfAdminReq(req); err != nil {
		return nil, err
	}

	res, err := s.auth.CheckIfAdmin(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return converter.CheckIfAdminResToPb(res), nil
}
