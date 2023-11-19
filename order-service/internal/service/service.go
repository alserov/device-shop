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
		Devices:   make([]*pb.Device, 0, len(req.Devices)),
		CreatedAt: time.Now().UTC(),
	}

	if err := utils.FetchDevices(ctx, s.deviceAddr, req.Devices, order); err != nil {
		return &pb.CreateOrderRes{}, err
	}

	if err := utils.ChangeBalance(ctx, s.userAddr, order); err != nil {
		utils.RollbackDeviceAmountPB(order.Devices, s.deviceAddr)
		return &pb.CreateOrderRes{}, err
	}

	if err := s.db.CreateOrder(ctx, order); err != nil {
		utils.RollbackDeviceAmountPB(order.Devices, s.deviceAddr)
		utils.RollBackBalance(order.UserUUID, utils.CountOrderPrice(order.Devices), s.userAddr)
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

	var (
		devices = make([]*pb.Device, 0, len(order.Devices))
		price   = float32(0)
		wg      = &sync.WaitGroup{}
		mu      = &sync.Mutex{}
		chErr   = make(chan error)
	)

	cl, cc, err := client.DialDevice(s.deviceAddr)
	if err != nil {
		return &pb.CheckOrderRes{}, err
	}
	defer cc.Close()

	wg.Add(len(order.Devices))

	for _, device := range order.Devices {
		device := device
		go func() {
			defer wg.Done()
			d, err := cl.GetDeviceByUUID(ctx, &pb.UUIDReq{UUID: device.DeviceUUID})
			if err != nil {
				chErr <- err
			}
			device := &pb.Device{
				UUID:         d.UUID,
				Title:        d.Title,
				Description:  d.Description,
				Price:        d.Price,
				Manufacturer: d.Manufacturer,
				Amount:       device.Amount,
			}

			mu.Lock()
			defer mu.Unlock()
			price += d.Price * float32(device.Amount)
			devices = append(devices, device)
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err = range chErr {
		return &pb.CheckOrderRes{}, err
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

// TODO: finish cancel order and fix incorrect order status

func (s service) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*emptypb.Empty, error) {
	if err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return &emptypb.Empty{}, err
	}
	if utils.StatusToCode(req.Status) == utils.CANCELED_CODE {
		order, err := s.db.CheckOrder(ctx, req.OrderUUID)
		if err != nil {
			return &emptypb.Empty{}, err
		}

		var (
			wg    = &sync.WaitGroup{}
			mu    = &sync.Mutex{}
			chErr = make(chan error)
			price = float32(0)
		)

		cl, cc, err := client.DialDevice(s.deviceAddr)
		if err != nil {
			return &emptypb.Empty{}, err
		}
		defer cc.Close()

		wg.Add(len(order.Devices))

		for _, device := range order.Devices {
			device := device
			go func() {
				defer wg.Done()
				d, err := cl.GetDeviceByUUID(ctx, &pb.UUIDReq{UUID: device.DeviceUUID})
				if err != nil {
					chErr <- err
				}
				mu.Lock()
				defer mu.Unlock()
				price += d.Price
			}()
		}

		go func() {
			wg.Wait()
			close(chErr)
		}()

		for err = range chErr {
			return &emptypb.Empty{}, err
		}
		utils.RollBackBalance(order.UserUUID, price, s.deviceAddr)
		utils.RollbackDeviceAmount(order.Devices, s.deviceAddr)
	}
	return &emptypb.Empty{}, nil
}
