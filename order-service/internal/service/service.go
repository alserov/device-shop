package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	txManager "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"

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

	conv *converter.ServiceConverter

	brokerAddr string

	txManager manager.TxManager
}

func NewService(ordersDB *sql.DB, broker *broker.Broker, log *slog.Logger) Service {
	return &service{
		log:        log,
		db:         postgres.NewOrderRepo(ordersDB),
		conv:       converter.NewServiceConverter(),
		brokerAddr: broker.BrokerAddr,
		txManager:  manager.NewTxManager(broker, log),
	}
}

func (s *service) CreateOrder(ctx context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	if err := s.txManager.DoTx(txManager.TxBody{
		OrderDevices: req.OrderDevices,
		OrderPrice:   req.OrderPrice,
		UserUUID:     req.UserUUID,
	}); err != nil {
		return models.CreateOrderRes{}, err
	}

	if err := s.db.CreateOrder(ctx, s.conv.CreateOrderReqToRepo(req, orderUUID)); err != nil {
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
	err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID)
	if err != nil {
		return err
	}

	return nil
}
