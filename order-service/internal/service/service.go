package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/order-service/internal/broker/consumer"
	"github.com/alserov/device-shop/order-service/internal/broker/producer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"

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
	orderTopic string
}

func NewService(ordersDB *sql.DB, brokerAddr string, orderTopic string, log *slog.Logger) Service {
	return &service{
		log:        log,
		db:         postgres.NewOrderRepo(ordersDB),
		conv:       converter.NewServiceConverter(),
		orderTopic: orderTopic,
		brokerAddr: brokerAddr,
	}
}

const (
	successStatus = 1
	internalError = "internal error"
	kafkaClientID = "ORDER_SERVICE"
)

func (s *service) CreateOrder(ctx context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	prod, err := producer.NewProducer([]string{s.brokerAddr}, kafkaClientID)
	if err != nil {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	_, _, err = prod.SendMessage(&sarama.ProducerMessage{
		Topic: s.orderTopic,
	})
	if err != nil {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	cons, err := sarama.NewConsumer([]string{s.brokerAddr}, &sarama.Config{})
	if err != nil {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	bytes, err := consumer.Subscribe(s.orderTopic, cons)
	if err != nil {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	txBytes := <-bytes

	var txRes models.TxResponse
	if err = json.Unmarshal(txBytes, &txRes); err != nil {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	if txRes.Status != successStatus {
		return models.CreateOrderRes{}, status.Error(codes.Internal, internalError)
	}

	if err = s.db.CreateOrder(ctx, s.conv.CreateOrderReqToRepo(req, orderUUID)); err != nil {
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
