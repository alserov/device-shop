package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/gateway/pkg/client"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/db/mongo"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
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
	cash, err := s.user.TopUpBalance(ctx, &pb.BalanceReq{
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
	cash, err := s.user.DebitBalance(ctx, &pb.BalanceReq{
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

func (s *service) AddToFavourite(ctx context.Context, req *pb.ChangeCollectionReq) (*emptypb.Empty, error) {
	cl, cc, err := client.DialDevice(s.deviceAddr)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer cc.Close()

	getDeviceReq := &pb.GetDeviceByUUIDReq{
		UUID: req.DeviceUUID,
	}

	dvc, err := cl.GetDeviceByUUID(ctx, getDeviceReq)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.collections.AddToFavourite(ctx, req.UserUUID, dvc); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) RemoveFromFavourite(ctx context.Context, req *pb.ChangeCollectionReq) (*emptypb.Empty, error) {
	err := s.collections.RemoveFromFavourite(ctx, &pb.ChangeCollectionReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetFavourite(ctx context.Context, req *pb.GetCollectionReq) (*pb.GetCollectionRes, error) {
	coll, err := s.collections.GetFavourite(ctx, req.UserUUID)
	if err != nil {
		return &pb.GetCollectionRes{}, err
	}

	var devices []*pb.Device

	for _, v := range coll {
		device := &pb.Device{
			UUID:         v.UUID,
			Title:        v.Title,
			Description:  v.Description,
			Price:        v.Price,
			Manufacturer: v.Manufacturer,
		}
		devices = append(devices, device)
	}

	return &pb.GetCollectionRes{
		Devices: devices,
	}, nil
}

func (s *service) AddToCart(ctx context.Context, req *pb.ChangeCollectionReq) (*emptypb.Empty, error) {
	cl, cc, err := client.DialDevice(s.deviceAddr)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	defer cc.Close()

	getDeviceReq := &pb.GetDeviceByUUIDReq{
		UUID: req.DeviceUUID,
	}

	dvc, err := cl.GetDeviceByUUID(ctx, getDeviceReq)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.collections.AddToCart(ctx, req.UserUUID, dvc); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) RemoveFromCart(ctx context.Context, req *pb.ChangeCollectionReq) (*emptypb.Empty, error) {
	err := s.collections.RemoveFromCart(ctx, &pb.ChangeCollectionReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.GetDeviceUUID(),
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetCart(ctx context.Context, req *pb.GetCollectionReq) (*pb.GetCollectionRes, error) {
	coll, err := s.collections.GetCart(ctx, req.UserUUID)
	if err != nil {
		return &pb.GetCollectionRes{}, err
	}

	var devices []*pb.Device

	for _, v := range coll {
		device := &pb.Device{
			UUID:         v.UUID,
			Title:        v.Title,
			Description:  v.Description,
			Price:        v.Price,
			Manufacturer: v.Manufacturer,
		}
		devices = append(devices, device)
	}

	return &pb.GetCollectionRes{
		Devices: devices,
	}, nil
}

func (s *service) RemoveDeviceFromCollections(ctx context.Context, req *pb.RemoveDeletedDeviceReq) (*emptypb.Empty, error) {
	if err := s.collections.RemoveDeviceFromCollections(ctx, req.DeviceUUID); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
