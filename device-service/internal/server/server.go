package server

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/device-service/internal/logger"
	"github.com/alserov/device-shop/device-service/internal/service"
	"github.com/alserov/device-shop/proto/gen/device"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func Register(s *grpc.Server, db *sql.DB, log *slog.Logger) {
	device.RegisterDevicesServer(s, &server{
		log:    log,
		device: service.NewService(db),
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
	device.UnimplementedDevicesServer
	log *slog.Logger
}

func (s *server) GetAllDevices(ctx context.Context, req *device.GetAllDevicesReq) (*device.DevicesRes, error) {
	devices, err := s.device.GetAllDevices(ctx, service.GetAllDevicesReq{
		Amount: req.GetAmount(),
		Index:  req.Index,
	})
	if err != nil {
		s.log.Error("failed to get all devices", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	var res device.DevicesRes
	for _, d := range devices {
		device := &device.Device{
			UUID:         d.UUID,
			Title:        d.Title,
			Description:  d.Description,
			Price:        d.Price,
			Manufacturer: d.Manufacturer,
			Amount:       d.Amount,
		}
		res.Devices = append(res.Devices, device)
	}

	return &res, nil
}

func (s *server) GetDevicesByTitle(ctx context.Context, req *device.GetDeviceByTitleReq) (*device.DevicesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetDevicesByManufacturer(ctx context.Context, manufacturer *device.GetByManufacturer) (*device.DevicesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetDevicesByPrice(ctx context.Context, price *device.GetByPrice) (*device.DevicesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetDeviceByUUID(ctx context.Context, req *device.GetDeviceByUUIDReq) (*device.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) mustEmbedUnimplementedDevicesServer() {
	//TODO implement me
	panic("implement me")
}
