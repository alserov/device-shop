package server

import (
	"context"
	"github.com/alserov/device-shop/proto/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Register(s *grpc.Server) {
	user.RegisterUsersServer(s, &server{})
}

type server struct {
	user.UnimplementedUsersServer
}

func (s server) TopUpBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DebitBalance(ctx context.Context, req *user.BalanceReq) (*user.BalanceRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) RemoveDeviceFromCollections(ctx context.Context, req *user.RemoveDeletedDeviceReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
