package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	txManager "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"github.com/alserov/device-shop/order-service/internal/utils/status"
	"log/slog"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error)
	CheckOrder(ctx context.Context, req models.CheckOrderReq) (models.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error
}

type service struct {
	log *slog.Logger

	db db.OrderRepo

	conv converter.ServiceConverter

	broker    *broker.Broker
	txManager manager.TxManager
}

func NewService(ordersDB *sql.DB, broker *broker.Broker, log *slog.Logger) Service {
	return &service{
		log:       log,
		db:        postgres.NewRepo(ordersDB, log),
		conv:      converter.NewServiceConverter(),
		broker:    broker,
		txManager: manager.NewTxManager(broker, log),
	}
}

func (s *service) CreateOrder(_ context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	err := s.txManager.CreateOrderTx(txManager.TxBody{
		OrderDevices: req.OrderDevices,
		OrderPrice:   req.OrderPrice,
		UserUUID:     req.UserUUID,
		Repo:         s.db,
		Order:        req,
		OrderUUID:    orderUUID,
	})
	if err != nil {
		return models.CreateOrderRes{}, err
	}

	return s.conv.CreateOrderResToService(orderUUID), nil
}

func (s *service) CheckOrder(ctx context.Context, req models.CheckOrderReq) (models.CheckOrderRes, error) {
	order, err := s.db.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return models.CheckOrderRes{}, err
	}

	return s.conv.CheckOrderToService(order), nil
}

func (s *service) UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error {
	if status.StatusToCode(req.Status) == status.CANCELED_CODE {

	}

	err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID)
	if err != nil {
		return err
	}

	return nil
}
