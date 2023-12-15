package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	broker "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	"github.com/alserov/device-shop/order-service/internal/db"
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"github.com/google/uuid"
	"log/slog"
	"sync"
)

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderReq) (string, error)
	CheckOrder(ctx context.Context, orderUUID string) (models.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error
	CancelOrder(ctx context.Context, orderUUID string) error
}

type service struct {
	log *slog.Logger

	repo db.OrderRepo

	conv converter.ServiceConverter

	txManager manager.TxManager
}

func NewService(repo db.OrderRepo, txManager manager.TxManager, log *slog.Logger) Service {
	return &service{
		log:       log,
		repo:      repo,
		conv:      converter.NewServiceConverter(),
		txManager: txManager,
	}
}

func (s *service) CreateOrder(_ context.Context, req models.CreateOrderReq) (string, error) {
	orderUUID := uuid.New().String()

	err := s.txManager.CreateOrderTx(broker.CreateOrderTxBody{
		OrderDevices: req.OrderDevices,
		OrderPrice:   req.OrderPrice,
		UserUUID:     req.UserUUID,
		Repo:         s.repo,
		Order:        req,
		OrderUUID:    orderUUID,
	})
	if err != nil {
		return "", err
	}

	return orderUUID, nil
}

func (s *service) CheckOrder(ctx context.Context, orderUUID string) (models.CheckOrderRes, error) {
	order, err := s.repo.CheckOrder(ctx, orderUUID)
	if err != nil {
		return models.CheckOrderRes{}, err
	}

	return s.conv.CheckOrderToService(order), nil
}

func (s *service) UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error {
	err := s.repo.UpdateOrder(ctx, req.Status, req.OrderUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CancelOrder(ctx context.Context, orderUUID string) error {
	var (
		wg           = &sync.WaitGroup{}
		chErr        = make(chan error)
		chTxs        = make(chan *sql.Tx, 2)
		orderDevices []repo.OrderDevice
		orderInfo    repo.CancelOrderRes
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		devices, tx, err := s.repo.CancelOrderDevicesTx(ctx, orderUUID)
		chTxs <- tx
		if err != nil {
			chErr <- err
		}
		orderDevices = devices
	}()

	go func() {
		defer wg.Done()
		info, tx, err := s.repo.CancelOrderTx(ctx, orderUUID)
		chTxs <- tx
		if err != nil {
			tx.Rollback()
			chErr <- err
		}
		orderInfo = info
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		for tx := range chTxs {
			tx.Rollback()
		}
		return err
	}

	err := s.txManager.CancelOrderTx(broker.CancelOrderTxBody{
		OrderUUID:    orderUUID,
		OrderDevices: orderDevices,
		OrderPrice:   orderInfo.Price,
		UserUUID:     orderInfo.UserUUID,
	})
	if err != nil {
		for tx := range chTxs {
			tx.Rollback()
		}
		return err
	}

	for tx := range chTxs {
		tx.Commit()
	}

	return nil
}
