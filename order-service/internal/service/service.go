package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service struct {
	db postgres.Repo
}

func New(db *sql.DB) pb.OrdersServer {
	return &service{
		db: postgres.New(db),
	}
}

func (s service) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderRes, error) {
	order := &postgres.CreateOrderReq{
		OrderUUID: uuid.New().String(),
		UserUUID:  req.UserUUID,
		Status:    utils.StatusToCode(utils.CREATING),
		Devices:   req.Devices,
	}

	if err := s.db.CreateOrder(ctx, order); err != nil {
		return &pb.CreateOrderRes{}, err
	}

	return &pb.CreateOrderRes{
		OrderUUID: order.OrderUUID,
	}, nil
}

func (s service) CheckOrder(ctx context.Context, req *pb.CheckOrderReq) (*pb.CheckOrderRes, error) {
	order, err := s.db.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return &pb.CheckOrderRes{}, err
	}

	devices := make([]*pb.Device, 0, len(order.Devices))
	price := float32(0)

	for _, d := range order.Devices {
		device := &pb.Device{
			UUID:         d.UUID,
			Title:        d.UUID,
			Description:  d.Description,
			Price:        d.Price,
			Manufacturer: d.Manufacturer,
			Amount:       d.Amount,
		}
		price += d.Price
		devices = append(devices, device)
	}

	return &pb.CheckOrderRes{
		Devices: devices,
		Status:  utils.StatusCodeToString(order.Status),
		Price:   price,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: order.CreatedAt.Unix(),
			Nanos:   int32(order.CreatedAt.Nanosecond()),
		},
	}, nil
}

func (s service) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*emptypb.Empty, error) {
	if err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
