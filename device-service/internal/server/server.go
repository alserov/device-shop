package server

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/device-service/internal/service"
	"github.com/alserov/device-shop/device-service/internal/utils/converter"
	"github.com/alserov/device-shop/device-service/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/proto/gen/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func Register(s *grpc.Server, db *sql.DB, log *slog.Logger, collectionAddr string) {
	device.RegisterDevicesServer(s, &server{
		log:     log,
		service: service.NewService(db, log),
		services: services{
			collectionAddr: collectionAddr,
		},
		valid: validation.NewValidator(),
		conv:  converter.NewServerConverter(),
	})
}

type server struct {
	log *slog.Logger

	device.UnimplementedDevicesServer
	service service.Service

	services services

	valid *validation.Validator
	conv  *converter.ServerConverter
}

type services struct {
	collectionAddr string
}

func (s *server) GetAllDevices(ctx context.Context, req *device.GetAllDevicesReq) (*device.DevicesRes, error) {
	devices, err := s.service.GetAllDevices(ctx, s.conv.Device.GetAllDevicesReqToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Device.GetAllDevicesResToPb(devices), nil
}

func (s *server) GetDevicesByTitle(ctx context.Context, req *device.GetDeviceByTitleReq) (*device.DevicesRes, error) {
	if err := s.valid.Device.ValidateGetDeviceByTitleReq(req); err != nil {
		return nil, err
	}

	devices, err := s.service.GetDevicesByTitle(ctx, req.Title)
	if err != nil {
		return nil, err
	}

	return s.conv.Device.GetAllDevicesResToPb(devices), nil
}

func (s *server) GetDevicesByManufacturer(ctx context.Context, req *device.GetByManufacturer) (*device.DevicesRes, error) {
	if err := s.valid.Device.ValidateGetDevicesByManufacturerReq(req); err != nil {
		return nil, err
	}

	devices, err := s.service.GetDevicesByManufacturer(ctx, req.Manufacturer)
	if err != nil {
		return nil, err
	}

	return s.conv.Device.GetAllDevicesResToPb(devices), err
}

func (s *server) GetDevicesByPrice(ctx context.Context, req *device.GetByPrice) (*device.DevicesRes, error) {
	if err := s.valid.Device.ValidateGetDevicesByPrice(req); err != nil {
		return nil, err
	}

	devices, err := s.service.GetDevicesByPrice(ctx, s.conv.Device.GetDevicesByPriceToService(req))
	if err != nil {
		return nil, err
	}

	return s.conv.Device.GetAllDevicesResToPb(devices), nil
}

func (s *server) GetDeviceByUUID(ctx context.Context, req *device.GetDeviceByUUIDReq) (*device.Device, error) {
	if err := s.valid.Device.ValidateGetDeviceByUUID(req); err != nil {
		return nil, err
	}

	device, err := s.service.GetDeviceByUUID(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Device.DeviceToPb(device), nil
}

func (s *server) CreateDevice(ctx context.Context, req *device.CreateDeviceReq) (*emptypb.Empty, error) {
	if err := s.valid.Admin.ValidateCreateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.service.CreateDevice(ctx, s.conv.Admin.CreateDeviceToService(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteDevice(ctx context.Context, req *device.DeleteDeviceReq) (*emptypb.Empty, error) {
	op := "server.DeleteDevice"

	if err := s.valid.Admin.ValidateDeleteDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.service.DeleteDevice(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	cl, cc, err := client.DialCollection(s.services.collectionAddr)
	if err != nil {
		s.log.Error("failed to dial collection service", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, "internal error")
	}
	defer cc.Close()

	_, err = cl.RemoveDeviceFromCollections(ctx, s.conv.Admin.RemoveDeviceFromCollectionsReqToPb(req))
	if err != nil {
		s.log.Error("failed to remove device from collections", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) UpdateDevice(ctx context.Context, req *device.UpdateDeviceReq) (*emptypb.Empty, error) {
	if err := s.valid.Admin.ValidateUpdateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.service.UpdateDevice(ctx, s.conv.Admin.UpdateDeviceToService(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) IncreaseDeviceAmount(ctx context.Context, req *device.IncreaseDeviceAmountByUUIDReq) (*emptypb.Empty, error) {
	if err := s.valid.Admin.ValidateIncreaseDeviceAmountReq(req); err != nil {
		return nil, err
	}

	err := s.service.IncreaseDeviceAmountByUUID(ctx, s.conv.Admin.IncreaseDeviceAmountToService(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
