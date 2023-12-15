package service

import (
	"context"
	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	broker "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"log/slog"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderReq) (string, error)
	CheckOrder(ctx context.Context, orderUUID string) (models.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error
	CancelOrder(ctx context.Context, orderUUID string, orderedDevices []models.OrderDevice) error
	GetOrderDevices(_ context.Context, orderUUID string) ([]models.OrderDevice, error)
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

func (s *service) CancelOrder(_ context.Context, orderUUID string, orderedDevices []models.OrderDevice) error {
	err := s.txManager.CancelOrderTx(broker.CancelOrderTxBody{
		Repo:         s.repo,
		OrderUUID:    orderUUID,
		OrderDevices: orderedDevices,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetOrderDeviceUUIDs(ctx context.Context, orderUUID string) ([]models.OrderDevice, error) {
	orderedDevices, err := s.repo.GetOrderDevices(ctx, orderUUID)
	if err != nil {
		return nil, err
	}
	return s.conv.OrderDevicesToService(orderedDevices), nil
}
