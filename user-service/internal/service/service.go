package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/gateway/pkg/client"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/db/mongo"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	conv "github.com/alserov/device-shop/user-service/internal/utils/converter"
	mg "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
)

type service struct {
	user        db.UserRepo
	collections db.CollectionsRepo
	deviceAddr  string
}

func New(pg *sql.DB, mg *mg.Client) pb.UsersServer {
	return &service{
		user:        postgres.NewRepo(pg),
		collections: mongo.NewRepo(mg),
		deviceAddr:  os.Getenv("DEVICE_ADDR"),
	}
}

func (s *service) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRes, error) {
	info, err := s.user.GetInfo(ctx, req.UserUUID)
	if err != nil {
		return &pb.GetUserInfoRes{}, err
	}

	return &pb.GetUserInfoRes{
		Cash:     info.Cash,
		Username: info.Username,
		Email:    info.Email,
		UUID:     info.UUID,
	}, nil
}

func (s *service) TopUpBalance(ctx context.Context, req *pb.BalanceReq) (*pb.BalanceRes, error) {
	cash, err := s.user.TopUpBalance(ctx, db.BalanceReq{
		Cash:     req.Cash,
		UserUUID: req.UserUUID,
	})
	if err != nil {
		return &pb.BalanceRes{}, err
	}

	return &pb.BalanceRes{
		Cash: cash,
	}, nil
}

func (s *service) DebitBalance(ctx context.Context, req *pb.BalanceReq) (*pb.BalanceRes, error) {
	cash, err := s.user.DebitBalance(ctx, db.BalanceReq{
		Cash:     req.Cash,
		UserUUID: req.UserUUID,
	})
	if err != nil {
		return &pb.BalanceRes{}, err
	}

	return &pb.BalanceRes{
		Cash: cash,
	}, nil
}
