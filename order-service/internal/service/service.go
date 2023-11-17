package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
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
		Devices:   make([]*pb.Device, 0, len(req.DeviceUUIDs)),
		CreatedAt: time.Now().UTC(),
	}

	var (
		wg        = &sync.WaitGroup{}
		chErr     = make(chan error)
		chDevices = make(chan *pb.Device, len(req.DeviceUUIDs))
	)
	wg.Add(len(req.DeviceUUIDs))
	{
		cl, cc, err := client.DialDevice(s.deviceAddr)
		if err != nil {
			chErr <- err
		}
		defer cc.Close()

		for _, v := range req.DeviceUUIDs {
			v := v
			go func() {
				defer wg.Done()
				device, err := cl.GetDeviceByUUID(ctx, &pb.UUIDReq{
					UUID: v,
				})
				if err != nil {
					chErr <- err
				}
				chDevices <- device
			}()
		}
	}

	go func() {
		wg.Wait()
		close(chDevices)
		close(chErr)
	}()

	go func() {
		for device := range chDevices {
			order.Devices = append(order.Devices, device)
		}
	}()

	for err := range chErr {
		return &pb.CreateOrderRes{}, err
	}

	tx, err := s.db.GetDB().Begin()
	if err != nil {
		return &pb.CreateOrderRes{}, err
	}
	chRPCErr := make(chan error)
	wg.Add(2)
	go func() {
		defer wg.Done()
		cl, cc, err := client.DialUser(s.userAddr)
		if err != nil {
			chRPCErr <- err
		}
		defer cc.Close()

		_, err = cl.DebitBalance(ctx, &pb.DebitBalanceReq{
			Cash:     float32(utils.CountOrderPrice(order.Devices)),
			UserUUID: order.UserUUID,
		})
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err = s.db.CreateOrder(ctx, order, tx); err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chRPCErr)
	}()
	for err = range chRPCErr {
		tx.Rollback()
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
