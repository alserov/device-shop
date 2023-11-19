package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/helpers"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"sync"
	"time"
)

type service struct {
	db         postgres.Repo
	deviceAddr string
	userAddr   string
}

func New(db *sql.DB) pb.OrdersServer {
	return &service{
		db:         postgres.New(db),
		deviceAddr: os.Getenv("DEVICE_ADDR"),
		userAddr:   os.Getenv("USER_ADDR"),
	}
}

func (s service) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderRes, error) {
	order := &entity.CreateOrderReqWithDevices{
		OrderUUID: uuid.New().String(),
		UserUUID:  req.UserUUID,
		Status:    utils.StatusToCode(utils.CREATING),
		Devices:   make([]*pb.Device, 0, len(req.Devices)),
		CreatedAt: time.Now().UTC(),
	}

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan *utils.RequestError, 2)
	)

	wg.Add(1)
	// RequestID = 1
	go helpers.FetchDevices(ctx, chErr, wg, s.deviceAddr, req.Devices, order.Devices)

	wg.Add(1)
	// RequestID = 2
	go helpers.ChangeBalance(ctx, chErr, wg, s.userAddr, order)

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		return &pb.CreateOrderRes{}, err.Handle(ctx, order, s.deviceAddr, s.userAddr)
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
