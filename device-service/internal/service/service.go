package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/db/mongo"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	mg "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net/http"
	"strings"
)

type service struct {
	postgres *sql.DB
	mongo    *mg.Client
}

func New(pg *sql.DB, mg *mg.Client) pb.DevicesServer {
	return &service{
		postgres: pg,
		mongo:    mg,
	}
}

func (s *service) CreateDevice(ctx context.Context, req *pb.CreateReq) (*emptypb.Empty, error) {
	r := &entity.Device{
		UUID:         uuid.New().String(),
		Title:        strings.ToLower(req.Title),
		Description:  req.Description,
		Price:        req.Price,
		Manufacturer: strings.ToLower(req.Manufacturer),
		Amount:       req.Amount,
	}

	if err := postgres.NewRepo(s.postgres).CreateDevice(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) DeleteDevice(ctx context.Context, req *pb.DeleteReq) (*emptypb.Empty, error) {

	if err := postgres.NewRepo(s.postgres).DeleteDevice(ctx, req.UUID); err != nil {
		return &emptypb.Empty{}, err
	}

	if err := mongo.NewRepo(s.mongo).DeleteWhereExists(ctx, req.UUID); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) UpdateDevice(ctx context.Context, req *pb.UpdateReq) (*emptypb.Empty, error) {
	r := &entity.UpdateDeviceReq{
		Title:       strings.ToLower(req.Title),
		Description: req.Description,
		Price:       req.Price,
		UUID:        req.UUID,
	}

	if err := postgres.NewRepo(s.postgres).UpdateDevice(ctx, r); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) GetAllDevices(ctx context.Context, req *pb.GetAllReq) (*pb.DevicesRes, error) {
	d, err := postgres.NewRepo(s.postgres).GetAllDevices(ctx, req.Index, req.Amount)
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

func (s *service) GetDevicesByTitle(ctx context.Context, req *pb.GetByTitleReq) (*pb.DevicesRes, error) {
	d, err := postgres.NewRepo(s.postgres).GetDevicesByTitle(ctx, strings.ToLower(req.Title))
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
		}
		devices = append(devices, pbDevice)
	}
	log.Println(devices, req.Title)
	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *service) GetDeviceByUUID(ctx context.Context, req *pb.UUIDReq) (*pb.Device, error) {
	d, err := postgres.NewRepo(s.postgres).GetDeviceByUUID(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	device := &pb.Device{
		UUID:         d.UUID,
		Title:        d.Title,
		Description:  d.Description,
		Manufacturer: d.Manufacturer,
		Price:        d.Price,
	}

	return device, nil
}

func (s *service) GetDevicesByManufacturer(ctx context.Context, req *pb.GetByManufacturer) (*pb.DevicesRes, error) {
	d, err := postgres.NewRepo(s.postgres).GetDevicesByManufacturer(ctx, req.Manufacturer)
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
		}
		devices = append(devices, pbDevice)
	}

	return &pb.DevicesRes{
		Devices: devices,
	}, nil
}

func (s *service) GetDevicesByPrice(ctx context.Context, req *pb.GetByPrice) (*pb.DevicesRes, error) {
	d, err := postgres.NewRepo(s.postgres).GetDevicesByPrice(ctx, uint(req.Min), uint(req.Max))
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
