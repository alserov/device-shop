package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/service"
	"github.com/alserov/device-shop/collection-service/internal/utils/converter"
	"github.com/alserov/device-shop/collection-service/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/proto/gen/collection"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type Server struct {
	GRPCServer *grpc.Server
	Repo       db.Repository
	Services   *Services
	Log        *slog.Logger
}

type Services struct {
	DeviceAddr string
}

func Register(s *Server) {
	collection.RegisterCollectionsServer(s.GRPCServer, &server{
		log:      s.Log,
		service:  service.NewService(s.Repo, s.Log),
		valid:    validation.NewValidator(),
		conv:     converter.NewServerConverter(),
		services: s.Services,
	})
}

type server struct {
	collection.UnsafeCollectionsServer
	service service.Service

	log *slog.Logger

	services *Services

	valid *validation.Validator
	conv  *converter.ServerConverter
}

const (
	internalErr = "internal error"
)

func (s *server) AddToFavourite(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	op := "server.AddToFavourite"

	if err := s.valid.Collection.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.services.DeviceAddr)
	if err != nil {
		s.log.Error("failed to dial device service", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalErr)
	}
	defer cc.Close()

	fetchedDevice, err := cl.GetDeviceByUUID(ctx, s.conv.Device.GetDeviceByUUIDReq(req.DeviceUUID))
	if err != nil {
		s.log.Error("failed get device by uuid", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalErr)
	}

	if err = s.service.AddToFavourite(ctx, req.UserUUID, s.conv.Device.PbDeviceToService(fetchedDevice)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) RemoveFromFavourite(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := s.valid.Collection.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	if err := s.service.RemoveFromFavourite(ctx, s.conv.Collection.ChangeCollectionReqToService(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetFavourite(ctx context.Context, req *collection.GetCollectionReq) (*collection.GetCollectionRes, error) {
	coll, err := s.service.GetFavourite(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.GetCollectionResToPb(coll), nil
}

func (s *server) AddToCart(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	op := "server.AddToCart"

	if err := s.valid.Collection.ValidateChangeCollectionReq(req); err != nil {
		return nil, err
	}

	cl, cc, err := client.DialDevice(s.services.DeviceAddr)
	if err != nil {
		s.log.Error("failed to dial device service", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalErr)
	}
	defer cc.Close()

	fetchedDevice, err := cl.GetDeviceByUUID(ctx, s.conv.Device.GetDeviceByUUIDReq(req.DeviceUUID))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.service.AddToCart(ctx, req.UserUUID, s.conv.Device.PbDeviceToService(fetchedDevice)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) RemoveFromCart(ctx context.Context, req *collection.ChangeCollectionReq) (*emptypb.Empty, error) {
	if err := s.service.RemoveFromCart(ctx, s.conv.Collection.ChangeCollectionReqToService(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetCart(ctx context.Context, req *collection.GetCollectionReq) (*collection.GetCollectionRes, error) {
	coll, err := s.service.GetCart(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.GetCollectionResToPb(coll), nil
}

func (s *server) RemoveDeviceFromCollections(ctx context.Context, req *collection.RemoveDeletedDeviceReq) (*emptypb.Empty, error) {
	if err := s.service.RemoveDeviceFromCollections(ctx, req.DeviceUUID); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
