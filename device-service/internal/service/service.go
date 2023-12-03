package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/db"
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/internal/service/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type service struct {
	db       db.DeviceRepo
	userAddr string
}

func NewService(db *sql.DB) Service {
	return &service{
		db: postgres.NewRepo(db),
	}
}

type Service interface {
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]*models.Device, error)
	GetAllDevices(ctx context.Context, req models.GetAllDevicesReq) ([]*models.Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]*models.Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error)
	GetDevicesByPrice(ctx context.Context, req models.GetByPrice) ([]*models.Device, error)

	CreateDevice(context.Context, models.Device) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, models.UpdateDeviceReq) error
	IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error
}

func (s *service) CreateDevice(ctx context.Context, req models.Device) error {
	err := s.db.CreateDevice(ctx, repo.Device{
		Title:        req.Title,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Description:  req.Description,
		Amount:       req.Amount,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteDevice(ctx context.Context, uuid string) error {
	if err := s.db.DeleteDevice(ctx, uuid); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateDevice(ctx context.Context, req models.UpdateDeviceReq) error {
	err := s.db.UpdateDevice(ctx, repo.UpdateDevice{
		UUID:        req.UUID,
		Description: req.Description,
		Title:       req.Title,
		Price:       req.Price,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllDevices(ctx context.Context, req models.GetAllDevicesReq) ([]*models.Device, error) {
	d, err := s.db.GetAllDevices(ctx, req.Index, req.Amount)
	if err != nil {
		return nil, err
	}

	devices := make([]*models.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &models.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return devices, nil
}

func (s *service) GetDevicesByTitle(ctx context.Context, title string) ([]*models.Device, error) {
	d, err := s.db.GetDevicesByTitle(ctx, strings.ToLower(title))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Nothing found by %s", title))
	}
	if err != nil {
		return nil, err
	}

	devices := make([]*models.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &models.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return devices, nil
}

func (s *service) GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error) {
	dv, err := s.db.GetDeviceByUUID(ctx, uuid)
	if err != nil {
		return models.Device{}, err
	}

	device := models.Device{
		UUID:         dv.UUID,
		Title:        dv.Title,
		Description:  dv.Description,
		Manufacturer: dv.Manufacturer,
		Price:        dv.Price,
		Amount:       dv.Amount,
	}

	return device, nil
}

func (s *service) GetDevicesByManufacturer(ctx context.Context, manu string) ([]*models.Device, error) {
	d, err := s.db.GetDevicesByManufacturer(ctx, manu)
	if errors.Is(err, sql.ErrNoRows) {
		return []*models.Device{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by %s", manu))
	}
	if err != nil {
		return []*models.Device{}, err
	}

	devices := make([]*models.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &models.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return devices, nil
}

func (s *service) GetDevicesByPrice(ctx context.Context, req models.GetByPrice) ([]*models.Device, error) {
	d, err := s.db.GetDevicesByPrice(ctx, uint(req.Min), uint(req.Max))
	if errors.Is(err, sql.ErrNoRows) {
		return []*models.Device{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by price > %d and price < %d", req.Min, req.Max))
	}
	if err != nil {
		return []*models.Device{}, err
	}

	devices := make([]*models.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &models.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
		}
		devices = append(devices, pbDevice)
	}

	return devices, nil
}

func (s *service) IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error {
	//TODO implement me
	panic("implement me")
}
