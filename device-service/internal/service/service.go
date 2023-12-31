package service

import (
	"context"
	"github.com/alserov/device-shop/device-service/internal/db"
	"github.com/alserov/device-shop/device-service/internal/service/models"
	"github.com/alserov/device-shop/device-service/internal/utils/converter"
	"github.com/google/uuid"
	"log/slog"
	"strings"
)

type service struct {
	log *slog.Logger

	repo db.Repository

	conv *converter.ServiceConverter
}

func NewService(repo db.Repository, log *slog.Logger) Service {
	return &service{
		log:  log,
		repo: repo,
		conv: converter.NewServiceConverter(),
	}
}

type Service interface {
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]models.Device, error)
	GetAllDevices(ctx context.Context, req models.GetAllDevicesReq) ([]models.Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]models.Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error)
	GetDevicesByPrice(ctx context.Context, req models.GetByPrice) ([]models.Device, error)

	CreateDevice(context.Context, models.CreateDeviceReq) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, models.UpdateDeviceReq) error
	IncreaseDeviceAmountByUUID(ctx context.Context, req models.IncreaseDeviceAmountReq) error
}

func (s *service) CreateDevice(ctx context.Context, req models.CreateDeviceReq) error {
	req.UUID = uuid.New().String()
	req.Title = strings.ToLower(req.Title)
	req.Manufacturer = strings.ToLower(req.Manufacturer)
	req.Description = strings.ToLower(req.Description)
	if err := s.repo.CreateDevice(ctx, s.conv.Admin.DeviceToRepo(req)); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteDevice(ctx context.Context, uuid string) error {
	if err := s.repo.DeleteDevice(ctx, uuid); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateDevice(ctx context.Context, req models.UpdateDeviceReq) error {
	if err := s.repo.UpdateDevice(ctx, s.conv.Admin.UpdateDeviceReqToRepo(req)); err != nil {
		return err
	}
	return nil
}

func (s *service) GetAllDevices(ctx context.Context, req models.GetAllDevicesReq) ([]models.Device, error) {
	devices, err := s.repo.GetAllDevices(ctx, req.Index, req.Amount)
	if err != nil {
		return nil, err
	}

	return s.conv.Device.DevicesToService(devices), nil
}

func (s *service) GetDevicesByTitle(ctx context.Context, title string) ([]models.Device, error) {
	devices, err := s.repo.GetDevicesByTitle(ctx, strings.ToLower(title))
	if err != nil {
		return nil, err
	}

	return s.conv.Device.DevicesToService(devices), nil
}

func (s *service) GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error) {
	foundDevice, err := s.repo.GetDeviceByUUID(ctx, uuid)
	if err != nil {
		return models.Device{}, err
	}

	return s.conv.Device.DeviceToService(foundDevice), nil
}

func (s *service) GetDevicesByManufacturer(ctx context.Context, manu string) ([]models.Device, error) {
	devices, err := s.repo.GetDevicesByManufacturer(ctx, manu)
	if err != nil {
		return nil, err
	}

	return s.conv.Device.DevicesToService(devices), nil
}

func (s *service) GetDevicesByPrice(ctx context.Context, req models.GetByPrice) ([]models.Device, error) {
	devices, err := s.repo.GetDevicesByPrice(ctx, uint(req.Min), uint(req.Max))
	if err != nil {
		return nil, err
	}

	return s.conv.Device.DevicesToService(devices), nil
}

func (s *service) IncreaseDeviceAmountByUUID(ctx context.Context, req models.IncreaseDeviceAmountReq) error {
	err := s.repo.IncreaseDeviceAmountByUUID(ctx, req.DeviceUUID, req.Amount)
	if err != nil {
		return err
	}

	return nil
}
