package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/entity"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"time"
)

type service struct {
	db postgres.Repository
}

func New(ordersDB, devicesDB, usersDB *sql.DB) pb.OrdersServer {
	return &service{
		db: postgres.New(ordersDB, devicesDB, usersDB),
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
		wg    *sync.WaitGroup
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

	return &pb.CreateOrderRes{
		OrderUUID: info.OrderUUID,
	}, nil
}

func (s service) CheckOrder(ctx context.Context, req *pb.CheckOrderReq) (*pb.CheckOrderRes, error) {
	order, err := s.db.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return &pb.CheckOrderRes{}, err
	}

	var (
		price = float32(0)
		wg    = &sync.WaitGroup{}
		mu    = &sync.Mutex{}
		chErr = make(chan error)
	)

	wg.Add(len(order.Devices))

	for _, d := range order.Devices {
		d := d
		go func() {
			defer wg.Done()
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
		return &pb.CheckOrderRes{}, err
	}

	return &pb.CheckOrderRes{
		Devices: order.Devices,
		Status:  utils.StatusCodeToString(order.Status),
		Price:   price,
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

		var (
			wg    = &sync.WaitGroup{}
			mu    = &sync.Mutex{}
			price = float32(0)
		)

		wg.Add(len(order.Devices))

		for _, device := range order.Devices {
			device := device
			go func() {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				price += device.Price
			}()
		}

		wg.Wait()
	}
	if err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return &pb.UpdateOrderRes{}, err
	}
	return &emptypb.Empty{}, nil
}
