package server

import (
	"context"
	"database/sql"

	"github.com/alserov/device-shop/proto/gen/user"
	"github.com/alserov/device-shop/user-service/internal/service"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"
	"github.com/alserov/device-shop/user-service/internal/utils/validation"

	"google.golang.org/grpc"

	"log/slog"
)

type Server struct {
	GRPCServer *grpc.Server
	DB         *sql.DB
	Log        *slog.Logger
	BrokerAddr string
	EmailTopic string
}

func Register(s *Server) {
	user.RegisterUsersServer(s.GRPCServer, &server{
		log:     s.Log,
		service: service.NewService(s.DB, s.Log, s.BrokerAddr, s.EmailTopic),
		valid:   validation.NewValidator(),
		conv:    converter.NewServerConverter(),
	})
}

type server struct {
	user.UnimplementedUsersServer
	service service.Service

	log   *slog.Logger
	valid *validation.Validator
	conv  *converter.ServerConverter
}

func (s *server) TopUpBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	if err := s.valid.Balance.ValidateBalanceReq(req); err != nil {
		return nil, err
	}

	balance, err := s.service.TopUpBalance(ctx, s.conv.Balance.BalanceReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Balance.BalanceResToPb(balance), nil
}

func (s *server) DebitBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	if err := s.valid.Balance.ValidateBalanceReq(req); err != nil {
		return nil, err
	}

	balance, err := s.service.DebitBalance(ctx, s.conv.Balance.BalanceReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Balance.BalanceResToPb(balance), nil
}

func (s *server) Signup(ctx context.Context, req *user.SignupReq) (*user.SignupRes, error) {
	if err := s.valid.Auth.ValidateSignupReq(req); err != nil {
		return nil, err
	}

	res, err := s.service.Signup(ctx, s.conv.Auth.SignupReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Auth.SignupResToPb(res), nil
}

func (s *server) Login(ctx context.Context, req *user.LoginReq) (*user.LoginRes, error) {
	if err := s.valid.Auth.ValidateLoginReq(req); err != nil {
		return nil, err
	}

	res, err := s.service.Login(ctx, s.conv.Auth.LoginReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Auth.LoginResToPb(res), nil
}

func (s *server) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoRes, error) {
	if err := s.valid.Info.ValidateGetUserInfoReq(req); err != nil {
		return nil, err
	}

	res, err := s.service.GetUserInfo(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Info.GetUserInfoResToPb(res), nil
}

func (s *server) CheckIfAdmin(ctx context.Context, req *user.CheckIfAdminReq) (*user.CheckIfAdminRes, error) {
	if err := s.valid.Info.ValidateCheckIfAdminReq(req); err != nil {
		return nil, err
	}

	res, err := s.service.CheckIfAdmin(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Info.CheckIfAdminResToPb(res), nil
}
