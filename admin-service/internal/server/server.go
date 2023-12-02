package server

import (
	"context"
	"database/sql"
	"github.com/alserov/admin-service/internal/logger"
	"github.com/alserov/admin-service/internal/service"
	"github.com/alserov/admin-service/internal/utils/converter"
	"github.com/alserov/admin-service/internal/utils/validation"
	"github.com/alserov/device-shop/proto/gen/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func Register(s *grpc.Server, db *sql.DB, log *slog.Logger) {
	admin.RegisterAdminServer(s, &server{
		admin: service.NewService(db),
		log:   log,
	})
}

type server struct {
	admin.UnimplementedAdminServer
	admin service.Service
	log   *slog.Logger
}

func (s *server) CreateDevice(ctx context.Context, req *admin.CreateDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateCreateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.admin.CreateDevice(ctx, converter.CreateDeviceToServiceStruct(req))
	if err != nil {
		s.log.Error("failed to create device: ", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteDevice(ctx context.Context, req *admin.DeleteDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateDeleteDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.admin.DeleteDevice(ctx, req.GetUUID())
	if err != nil {
		s.log.Error("failed to delete device: ", logger.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) UpdateDevice(ctx context.Context, req *admin.UpdateDeviceReq) (*emptypb.Empty, error) {
	if err := validation.ValidateUpdateDeviceReq(req); err != nil {
		return nil, err
	}

	err := s.admin.UpdateDevice(ctx, converter.UpdateDeviceToServiceStruct(req))
	if err != nil {
		s.log.Error("failed to update device: ", logger.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
