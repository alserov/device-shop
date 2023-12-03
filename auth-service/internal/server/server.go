package server

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/auth-service/internal/service"
	"github.com/alserov/device-shop/auth-service/internal/utils/converter"
	"github.com/alserov/device-shop/auth-service/internal/utils/validation"
	"github.com/alserov/device-shop/proto/gen/auth"
	"google.golang.org/grpc"
	"log/slog"
)

func Register(s *grpc.Server, db *sql.DB, log *slog.Logger, kafkaTopic string, kafkaBrokerAddr string) {
	auth.RegisterAuthServer(s, &server{
		auth: service.NewService(db, kafkaTopic, kafkaBrokerAddr),
		log:  log,
	})
}

type server struct {
	auth.UnimplementedAuthServer
	auth service.Service
	log  *slog.Logger
}

func (s server) Signup(ctx context.Context, req *auth.SignupReq) (*auth.SignupRes, error) {
	if err := validation.ValidateSignupReq(req); err != nil {
		return nil, err
	}

	res, err := s.auth.Signup(ctx, converter.SignupReqToServiceStruct(req))
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
