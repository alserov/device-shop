package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/entity"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"time"
)

type service struct {
	deviceAddr string
	db         postgres.Repository
}

func New(ordersDB, devicesDB, usersDB *sql.DB, deviceAddr string) pb.OrdersServer {
	return &service{
		deviceAddr: deviceAddr,
		db:         postgres.New(ordersDB, devicesDB, usersDB),
	}
}

func (s service) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderRes, error) {
	var (
		err  error
		now  = time.Now().UTC()
		info = &entity.OrderAdditional{
			CreatedAt: &now,
			Status:    utils.CREATING_CODE,
			OrderUUID: uuid.New().String(),
		}
		wg    = &sync.WaitGroup{}
		chErr = make(chan error)
	)

	info.TotalPrice, err = utils.CountPrice(ctx, req.Devices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	orderTx, err := s.db.GetOrdersDB().Begin()
	if err != nil {
		return &pb.CreateOrderRes{}, err
	}

	deviceTx, err := s.db.GetDevicesDB().Begin()
	if err != nil {
		return &pb.CreateOrderRes{}, err
	}

	balanceTx, err := s.db.GetUsersDB().Begin()
	if err != nil {
		return &pb.CreateOrderRes{}, err
	}

	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := s.db.CreateOrder(ctx, orderTx, req, info); err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err = s.db.DecreaseDevicesAmount(ctx, deviceTx, req.Devices); err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err = s.db.DebitBalance(ctx, balanceTx, req.UserUUID, info.TotalPrice); err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err = range chErr {
		cancel()
		orderTx.Rollback()
		balanceTx.Rollback()
		deviceTx.Rollback()
		return &pb.CreateOrderRes{}, err
	}

	orderTx.Commit()
	balanceTx.Commit()
	deviceTx.Commit()

	return &pb.CreateOrderRes{
		OrderUUID: info.OrderUUID,
	}, nil
}

func (s service) CheckOrder(ctx context.Context, req *pb.CheckOrderReq) (*pb.CheckOrderRes, error) {
	order, err := s.db.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return &pb.CheckOrderRes{}, err
	}

	cl, cc, err := client.DialDevice(s.deviceAddr)
	if err != nil {
		return &pb.CheckOrderRes{}, err
	}
	defer cc.Close()

	var (
		wg        = &sync.WaitGroup{}
		chDevices = make(chan *pb.Device, len(order.Devices))
		chErr     = make(chan error)
		devices   = make([]*pb.Device, 0, len(order.Devices))
	)
	wg.Add(len(order.Devices))
	for _, d := range order.Devices {
		d := d
		go func() {
			defer wg.Done()
			device, err := cl.GetDeviceByUUID(ctx, &pb.GetDeviceByUUIDReq{
				UUID: d.UUID,
			})
			if err != nil {
				chErr <- err
			}
			device.Amount = d.Amount
			chDevices <- device
		}()
	}

	go func() {
		wg.Wait()
		close(chDevices)
		close(chErr)
	}()

	for err = range chErr {
		return &pb.CheckOrderRes{}, err
	}

	for device := range chDevices {
		devices = append(devices, device)
	}

	return &pb.CheckOrderRes{
		Devices: devices,
		Status:  utils.StatusCodeToString(order.Status),
		Price:   order.TotalPrice,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: order.CreatedAt.Unix(),
			Nanos:   int32(order.CreatedAt.Nanosecond()),
		},
	}, nil
}

func (s service) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*pb.UpdateOrderRes, error) {
	if utils.StatusToCode(req.Status) == utils.CANCELED_CODE {
		order, err := s.db.CheckOrder(ctx, req.OrderUUID)
		if err != nil {
			return &pb.UpdateOrderRes{}, err
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		deviceTx, err := s.db.GetDevicesDB().Begin()
		if err != nil {
			return &pb.UpdateOrderRes{}, err
		}

		balanceTx, err := s.db.GetUsersDB().Begin()
		if err != nil {
			return &pb.UpdateOrderRes{}, err
		}

		var (
			chErr = make(chan error)
			wg    = &sync.WaitGroup{}
		)
		wg.Add(2)

		go func() {
			defer wg.Done()
			if err = s.db.RollbackDevices(ctx, deviceTx, order.Devices); err != nil {
				chErr <- err
			}
		}()

		go func() {
			defer wg.Done()
			if err = s.db.RollbackBalance(ctx, balanceTx, order.UserUUID, order.UUID, order.TotalPrice); err != nil {
				chErr <- err
			}
		}()

		go func() {
			wg.Wait()
			close(chErr)
		}()

		for err = range chErr {
			return &pb.UpdateOrderRes{}, err
		}

		balanceTx.Commit()
		deviceTx.Commit()
	}
	if err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return &pb.UpdateOrderRes{}, err
	}
	return &pb.UpdateOrderRes{}, nil
}
