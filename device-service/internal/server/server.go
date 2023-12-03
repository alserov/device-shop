package server

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/device-service/internal/logger"
	"github.com/alserov/device-shop/device-service/internal/service"
	"github.com/alserov/device-shop/device-service/internal/service/models"
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
		log:    log,
		device: service.NewService(db),
		services: ServicesAddr{
			collectionAddr: collectionAddr,
		},
	})
}

type Server interface {
	GetAllDevices(context.Context, *device.GetAllDevicesReq) (*device.DevicesRes, error)
	GetDevicesByTitle(context.Context, *device.GetDeviceByTitleReq) (*device.DevicesRes, error)
	GetDevicesByManufacturer(context.Context, *device.GetByManufacturer) (*device.DevicesRes, error)
	GetDevicesByPrice(context.Context, *device.GetByPrice) (*device.DevicesRes, error)
	GetDeviceByUUID(context.Context, *device.GetDeviceByUUIDReq) (*device.Device, error)
}

type server struct {
	device service.Service
	log    *slog.Logger

	services ServicesAddr

	device.UnimplementedDevicesServer
}

type ServicesAddr struct {
	collectionAddr string
}

func (s *server) GetAllDevices(ctx context.Context, req *device.GetAllDevicesReq) (*device.DevicesRes, error) {
	devices, err := s.device.GetAllDevices(ctx, models.GetAllDevicesReq{
		Amount: req.GetAmount(),
		Index:  req.Index,
	})
	if err != nil {
		s.log.Error("failed to get all devices", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	var res device.DevicesRes
	for _, d := range devices {
		device := converter.DeviceToPb(*d)
		res.Devices = append(res.Devices, device)
	}

	return &res, nil
}

func (s *server) GetDevicesByTitle(ctx context.Context, req *device.GetDeviceByTitleReq) (*device.DevicesRes, error) {
	if err := validation.ValidateGetDeviceByTitleReq(req); err != nil {
		return nil, err
	}

	foundDevices, err := s.device.GetDevicesByTitle(ctx, req.Title)
	if err != nil {
		return nil, err
	}

	var devices []*device.Device
	for _, d := range foundDevices {
		device := converter.DeviceToPb(*d)
		devices = append(devices, device)
	}

	return &device.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *server) GetDevicesByManufacturer(ctx context.Context, req *device.GetByManufacturer) (*device.DevicesRes, error) {
	if err := validation.ValidateGetDevicesByManufacturerReq(req); err != nil {
		return nil, err
	}

	foundDevices, err := s.device.GetDevicesByManufacturer(ctx, req.Manufacturer)
	if err != nil {
		return nil, err
	}

	var devices []*device.Device
	for _, d := range foundDevices {
		device := converter.DeviceToPb(*d)
		devices = append(devices, device)
	}

	return &device.DevicesRes{
		Devices: devices,
	}, err
}

func (s *server) GetDevicesByPrice(ctx context.Context, req *device.GetByPrice) (*device.DevicesRes, error) {
	if err := validation.ValidateGetDevicesByPrice(req); err != nil {
		return nil, err
	}

	foundDevices, er := s.device.GetDevicesByPrice(ctx, converter.GetDevicesByPriceToRepo(req))
}

func (s *server) GetDeviceByUUID(ctx context.Context, req *device.GetDeviceByUUIDReq) (*device.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) CreateDevice(ctx context.Context, req *device.CreateDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateCreateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.device.CreateDevice(ctx, converter.CreateDeviceToService(req))
	if err != nil {
		s.log.Error("failed to create device: ", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteDevice(ctx context.Context, req *device.DeleteDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateDeleteDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.device.DeleteDevice(ctx, req.UUID)
	if err != nil {
		s.log.Error("failed to delete device: ", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	cl, cc, err := client.DialCollection(s.services.collectionAddr)
	if err != nil {
		s.log.Error("failed to dial collection service: " + err.Error())
		return nil, status.Error(codes.Internal, "internal error")
	}
	defer cc.Close()

	return &emptypb.Empty{}, nil
}

func (s *server) UpdateDevice(ctx context.Context, req *device.UpdateDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateUpdateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.device.UpdateDevice(ctx, converter.UpdateDeviceToService(req))
	if err != nil {
		s.log.Error("failed to update device: ", logger.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
