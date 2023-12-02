package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/db"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
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
	GetDevicesByManufacturer(ctx context.Context, req GetByManufacturer) ([]*Device, error)
	GetAllDevices(ctx context.Context, req GetAllDevicesReq) ([]*Device, error)
	GetDevicesByTitle(ctx context.Context, req GetDeviceByTitleReq) ([]*Device, error)
	GetDeviceByUUID(ctx context.Context, req GetDeviceByUUIDReq) (Device, error)
	GetDevicesByPrice(ctx context.Context, req GetByPrice) ([]*Device, error)
}

func (s *service) GetAllDevices(ctx context.Context, req GetAllDevicesReq) ([]*Device, error) {
	d, err := s.db.GetAllDevices(ctx, req.Index, req.Amount)
	if err != nil {
		return nil, err
	}

	devices := make([]*Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &Device{
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

func (s *service) GetDevicesByTitle(ctx context.Context, req GetDeviceByTitleReq) ([]*Device, error) {
	d, err := s.db.GetDevicesByTitle(ctx, strings.ToLower(req.Title))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by %s", req.Title))
	}
	if err != nil {
		return nil, err
	}

	devices := make([]*Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &Device{
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

func (s *service) GetDeviceByUUID(ctx context.Context, req GetDeviceByUUIDReq) (Device, error) {
	dv, err := s.db.GetDeviceByUUID(ctx, req.UUID)
	if err != nil {
		return Device{}, err
	}

	device := Device{
		UUID:         dv.UUID,
		Title:        dv.Title,
		Description:  dv.Description,
		Manufacturer: dv.Manufacturer,
		Price:        dv.Price,
		Amount:       dv.Amount,
	}

	return device, nil
}

func (s *service) GetDevicesByManufacturer(ctx context.Context, req GetByManufacturer) ([]*Device, error) {
	d, err := s.db.GetDevicesByManufacturer(ctx, req.Manufacturer)
	if errors.Is(err, sql.ErrNoRows) {
		return []*Device{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by %s", req.Manufacturer))
	}
	if err != nil {
		return []*Device{}, err
	}

	devices := make([]*Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &Device{
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

func (s *service) GetDevicesByPrice(ctx context.Context, req GetByPrice) ([]*Device, error) {
	d, err := s.db.GetDevicesByPrice(ctx, uint(req.Min), uint(req.Max))
	if errors.Is(err, sql.ErrNoRows) {
		return []*Device{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by price > %d and price < %d", req.Min, req.Max))
	}
	if err != nil {
		return []*Device{}, err
	}

	devices := make([]*Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &Device{
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
