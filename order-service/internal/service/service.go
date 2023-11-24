package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/db"
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
	order      db.OrderRepo
	user       db.UserRepo
	device     db.DeviceRepo
}

func New(ordersDB, devicesDB, usersDB *sql.DB, deviceAddr string) pb.OrdersServer {
	return &service{
		deviceAddr: deviceAddr,
		order:      postgres.NewOrderRepo(ordersDB),
		device:     postgres.NewDeviceRepo(devicesDB),
		user:       postgres.NewUserRepo(usersDB),
	}
}

const (
	_ = iota
	txOne
	txTwo
	txThree
)

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
		txCh  = make(chan *sql.Tx, txThree)
	)

	info.TotalPrice, err = utils.CountPrice(ctx, req.Devices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(txThree)

	go func() {
		defer wg.Done()
		if err := s.order.CreateOrder(ctx, txCh, req, info); err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.device.DecreaseDevicesAmount(ctx, txCh, req.Devices); err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.user.DebitBalance(ctx, txCh, req.UserUUID, info.TotalPrice); err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
		close(txCh)
	}()

	for err = range chErr {
		cancel()
		for tx := range txCh {
			tx.Rollback()
		}
		return &pb.CreateOrderRes{}, err
	}

	for tx := range txCh {
		tx.Commit()
	}

	return &pb.CreateOrderRes{
		OrderUUID: info.OrderUUID,
	}, nil
}

func (s service) CheckOrder(ctx context.Context, req *pb.CheckOrderReq) (*pb.CheckOrderRes, error) {
	order, err := s.order.CheckOrder(ctx, req.OrderUUID)
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
		order, err := s.order.CheckOrder(ctx, req.OrderUUID)
		if err != nil {
			return &pb.UpdateOrderRes{}, err
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		var (
			chErr = make(chan error)
			wg    = &sync.WaitGroup{}
		)
		wg.Add(txTwo)
		txCh := make(chan *sql.Tx, txTwo)

		go func() {
			defer wg.Done()
			if err = s.device.RollbackDevices(ctx, txCh, order.Devices); err != nil {
				chErr <- err
			}
		}()

		go func() {
			defer wg.Done()
			o, err := s.order.CheckOrder(ctx, req.OrderUUID)
			if o.Status == utils.CANCELED_CODE {
				chErr <- errors.New("order is already canceled")
				return
			}
			if err = s.user.RollbackBalance(ctx, txCh, order.UserUUID, order.UUID, order.TotalPrice); err != nil {
				chErr <- err
			}
		}()

		go func() {
			wg.Wait()
			close(chErr)
			close(txCh)
		}()

		for err = range chErr {
			for tx := range txCh {
				tx.Rollback()
			}
			return &pb.UpdateOrderRes{}, err
		}

		for tx := range txCh {
			tx.Commit()
		}
	}
	if err := s.order.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return &pb.UpdateOrderRes{}, err
	}
	return &pb.UpdateOrderRes{}, nil
}
