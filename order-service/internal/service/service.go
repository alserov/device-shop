package service

import (
	"context"

	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	broker "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	"github.com/alserov/device-shop/order-service/internal/db"
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

func (s *service) CreateOrder(_ context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	err := s.txManager.CreateOrderTx(broker.TxBody{
		OrderDevices: req.OrderDevices,
		OrderPrice:   req.OrderPrice,
		UserUUID:     req.UserUUID,
		Repo:         s.repo,
		Order:        req,
		OrderUUID:    orderUUID,
	})
	if err != nil {
		return models.CreateOrderRes{}, err
	}

	return s.conv.CreateOrderResToService(orderUUID), nil
}

func (s *service) CheckOrder(ctx context.Context, req models.CheckOrderReq) (models.CheckOrderRes, error) {
	order, err := s.repo.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return models.CheckOrderRes{}, err
	}

	return s.conv.CheckOrderToService(order), nil
}

// TODO: finish update order func
func (s *service) UpdateOrder(ctx context.Context, req models.UpdateOrderReq) error {
	if status.StatusToCode(req.Status) == status.CANCELED_CODE {

	}

	err := s.repo.UpdateOrder(ctx, req.Status, req.OrderUUID)
	if err != nil {
		return err
	}

	return nil
}
