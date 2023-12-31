package server

import (
	"context"
	"github.com/alserov/device-shop/order-service/internal/db"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/broker/manager"
	"github.com/alserov/device-shop/order-service/internal/service"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"github.com/alserov/device-shop/order-service/internal/utils/validation"

	"github.com/alserov/device-shop/proto/gen/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Server struct {
	Log *slog.Logger

	Services   *Services
	GRPCServer *grpc.Server
	Repo       db.Repository

	Broker *broker.Broker
}

type Services struct {
	DeviceAddr string
}

func Register(s *Server) {
	manager := manager.NewTxManager(s.Broker, s.Log)

	order.RegisterOrdersServer(s.GRPCServer, &server{
		log:     s.Log,
		service: service.NewService(s.Repo, manager, s.Log),
		valid:   validation.NewValidator(),
		conv:    converter.NewServerConverter(),
		services: services{
			deviceAddr: s.Services.DeviceAddr,
		},
	})
}

type server struct {
	log *slog.Logger

	order.UnimplementedOrdersServer
	service service.Service

	services services

	conv  converter.ServerConverter
	valid validation.Validator
}

type services struct {
	deviceAddr string
}

const (
	internalError = "internal error"
)

func (s *server) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderRes, error) {
	op := "server.CreateOrder"
	if err := s.valid.ValidateCreateOrderReq(req); err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.services.deviceAddr)
	if err != nil {
		s.log.Error("failed to dial device service", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}
	defer cc.Close()

	orderPrice, err := utils.CountOrderPrice(ctx, cl, req.OrderDevices)
	if err != nil {
		return nil, err
	}

	orderUUID, err := s.service.CreateOrder(ctx, s.conv.CreateOrderReqToService(req, orderPrice))
	if err != nil {
		return nil, err
	}

	return s.conv.CreateOrderResToPb(orderUUID), nil
}

func (s *server) CheckOrder(ctx context.Context, req *order.CheckOrderReq) (*order.CheckOrderRes, error) {
	op := "server.CheckOrder"

	if err := s.valid.ValidateCheckOrderReq(req); err != nil {
		return nil, err
	}

	order, err := s.service.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.services.deviceAddr)
	if err != nil {
		s.log.Error("failed to dial device service", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}
	defer cc.Close()

	devicesFromOrder, err := utils.FetchDevicesFromOrder(ctx, cl, order.OrderDevices)
	if err != nil {
		s.log.Error("failed to fetch devices from order", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}

	return s.conv.CheckOrderResToPb(order, devicesFromOrder), nil
}

func (s *server) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (*order.UpdateOrderRes, error) {
	if err := s.valid.ValidateUpdateOrderReq(req); err != nil {
		return nil, err
	}

	err := s.service.UpdateOrder(ctx, s.conv.UpdateOrderReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.UpdateOrderResToPb(req.Status), nil
}

func (s *server) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (*emptypb.Empty, error) {
	if err := s.valid.ValidateCancelOrderReq(req); err != nil {
		return nil, err
	}

	err := s.service.CancelOrder(ctx, req.OrderUUID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
