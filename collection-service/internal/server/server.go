package server

import (
	"context"
	"github.com/alserov/device-shop/collection-service/internal/service"
	"github.com/alserov/device-shop/collection-service/internal/utils/converter"
	"github.com/alserov/device-shop/collection-service/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/alserov/device-shop/proto/gen/device"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func Register(s *grpc.Server, db *mongo.Client, log *slog.Logger) {
	collection.RegisterCollectionsServer(s, &server{
		collections: service.NewService(db),
		log:         log,
	})
}

type server struct {
	collection.UnsafeCollectionsServer
	collections service.Service

	log *slog.Logger

	deviceServiceAddr string
}

func (s *server) AddToFavourite(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := validation.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.deviceServiceAddr)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	fetchedDevice, err := cl.GetDeviceByUUID(ctx, &device.GetDeviceByUUIDReq{
		UUID: req.DeviceUUID,
	})
	if err != nil {
		return nil, err
	}

	if err = s.collections.AddToFavourite(ctx, req.UserUUID, converter.PbDeviceToServiceStruct(fetchedDevice)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) RemoveFromFavourite(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := validation.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	if err := s.collections.RemoveFromFavourite(ctx, converter.PbChangeCollectionReqTpServiceStruct(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetFavourite(ctx context.Context, req *collection.GetCollectionReq) (*collection.GetCollectionRes, error) {
	coll, err := s.collections.GetFavourite(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	var devices []*device.Device
	for _, v := range coll {
		device := converter.ServiceDeviceToPb(*v)
		devices = append(devices, device)
	}

	return &collection.GetCollectionRes{
		Devices: devices,
	}, nil
}

func (s *server) AddToCart(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := validation.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.deviceServiceAddr)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	fetchedDevice, err := cl.GetDeviceByUUID(ctx, req.DeviceUUID)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.collections.AddToCart(ctx, req.UserUUID, converter.PbDeviceToServiceStruct(fetchedDevice)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) RemoveFromCart(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := s.collections.RemoveFromCart(ctx, converter.PbChangeCollectionReqTpServiceStruct(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetCart(ctx context.Context, req *collection.GetCollectionReq) (*collection.GetCollectionRes, error) {
	coll, err := s.collections.GetCart(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	var devices []*device.Device

	for _, v := range coll {
		device := converter.ServiceDeviceToPb(*v)
		devices = append(devices, device)
	}

	return &collection.GetCollectionRes{
		Devices: devices,
	}, nil
}

func (s *server) RemoveDeviceFromCollections(ctx context.Context, req *collection.RemoveDeletedDeviceReq) (*emptypb.Empty, error) {
	if err := s.collections.RemoveDeviceFromCollections(ctx, req.DeviceUUID); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
