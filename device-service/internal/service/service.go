package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"os"
	"strings"
)

type service struct {
	postgres postgres.Repository
	userAddr string
}

func New(pg *sql.DB) pb.DevicesServer {
	return &service{
		postgres: postgres.NewRepo(pg),
		userAddr: os.Getenv("USER_ADDR"),
	}
}

func (s *service) CreateDevice(ctx context.Context, req *pb.CreateDeviceReq) (*emptypb.Empty, error) {
	r := &entity.Device{
		UUID:         uuid.New().String(),
		Title:        strings.ToLower(req.Title),
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: strings.ToLower(req.Manufacturer),
		Amount:       req.Amount,
	}

	if err := s.postgres.CreateDevice(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) UpdateDevice(ctx context.Context, req *pb.UpdateDeviceReq) (*emptypb.Empty, error) {
	r := &entity.UpdateDeviceReq{
		Title:       strings.ToLower(req.Title),
		Description: req.Description,
		Price:       req.Price,
		UUID:        req.UUID,
	}

	if err := s.postgres.UpdateDevice(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceReq) (*emptypb.Empty, error) {
	chErr := make(chan error)
	go func() {
		defer close(chErr)
		cl, cc, err := client.DialUser(s.userAddr)
		if err != nil {
			chErr <- err
		}
		defer cc.Close()

		_, err = cl.RemoveDeviceFromCollections(ctx, &pb.RemoveDeletedDeviceReq{
			DeviceUUID: req.UUID,
		})
		if err != nil {
			chErr <- err
		}
	}()

	if err := s.postgres.DeleteDevice(ctx, req.UUID); err != nil {
		return &emptypb.Empty{}, err
	}

	for err := range chErr {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetAllDevices(ctx context.Context, req *pb.GetAllDevicesReq) (*pb.DevicesRes, error) {
	d, err := s.postgres.GetAllDevices(ctx, req.Index, req.Amount)
	if err != nil {
		return &pb.DevicesRes{}, err
	}

	devices := make([]*pb.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &pb.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *service) GetDevicesByTitle(ctx context.Context, req *pb.GetDeviceByTitleReq) (*pb.DevicesRes, error) {
	d, err := s.postgres.GetDevicesByTitle(ctx, strings.ToLower(req.Title))
	if errors.Is(err, sql.ErrNoRows) {
		return &pb.DevicesRes{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by %s", req.Title))
	}
	if err != nil {
		return &pb.DevicesRes{}, err
	}

	devices := make([]*pb.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &pb.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *service) GetDeviceByUUID(ctx context.Context, req *pb.GetDeviceByUUIDReq) (*pb.Device, error) {
	dv, err := s.postgres.GetDeviceByUUID(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	device := &pb.Device{
		UUID:         dv.UUID,
		Title:        dv.Title,
		Description:  dv.Description,
		Manufacturer: dv.Manufacturer,
		Price:        dv.Price,
		Amount:       dv.Amount,
	}

	return device, nil
}

func (s *service) GetDevicesByManufacturer(ctx context.Context, req *pb.GetByManufacturer) (*pb.DevicesRes, error) {
	d, err := s.postgres.GetDevicesByManufacturer(ctx, req.Manufacturer)
	if errors.Is(err, sql.ErrNoRows) {
		return &pb.DevicesRes{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by %s", req.Manufacturer))
	}
	if err != nil {
		return &pb.DevicesRes{}, err
	}

	devices := make([]*pb.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &pb.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
			Amount:       dv.Amount,
		}
		devices = append(devices, pbDevice)
	}

	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *service) GetDevicesByPrice(ctx context.Context, req *pb.GetByPrice) (*pb.DevicesRes, error) {
	d, err := s.postgres.GetDevicesByPrice(ctx, uint(req.Min), uint(req.Max))
	if errors.Is(err, sql.ErrNoRows) {
		return &pb.DevicesRes{}, status.Error(http.StatusBadRequest, fmt.Sprintf("Nothing found by price > %d and price < %d", req.Min, req.Max))
	}
	if err != nil {
		return &pb.DevicesRes{}, err
	}

	devices := make([]*pb.Device, 0, len(d))
	for _, dv := range d {
		pbDevice := &pb.Device{
			UUID:         dv.UUID,
			Title:        dv.Title,
			Description:  dv.Description,
			Manufacturer: dv.Manufacturer,
			Price:        dv.Price,
		}
		devices = append(devices, pbDevice)
	}

	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}
