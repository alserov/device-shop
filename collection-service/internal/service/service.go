package service

import (
	"context"

	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/service/models"
	"github.com/alserov/device-shop/collection-service/internal/utils/converter"

	"log/slog"
)

type Service interface {
	AddToFavourite(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromFavourite(ctx context.Context, req models.ChangeCollectionReq) error
	GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error)

	AddToCart(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromCart(ctx context.Context, req models.ChangeCollectionReq) error
	GetCart(ctx context.Context, userUUID string) ([]*models.Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}

func NewService(repo db.Repository, log *slog.Logger) Service {
	return &service{
		log:  log,
		repo: repo,
		conv: converter.NewServiceConverter(),
	}
}

type service struct {
	log  *slog.Logger
	repo db.Repository

	conv *converter.ServiceConverter
}

func (s *service) AddToFavourite(ctx context.Context, userUUID string, device models.Device) error {
	if err := s.repo.AddToFavourite(ctx, userUUID, s.conv.Device.DeviceToRepo(device)); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveFromFavourite(ctx context.Context, req models.ChangeCollectionReq) error {
	if err := s.repo.RemoveFromFavourite(ctx, req.UserUUID, req.DeviceUUID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error) {
	col, err := s.repo.GetFavourite(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.CollectionToServer(col), nil
}

func (s *service) AddToCart(ctx context.Context, userUUID string, device models.Device) error {
	if err := s.repo.AddToCart(ctx, userUUID, s.conv.Device.DeviceToRepo(device)); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveFromCart(ctx context.Context, req models.ChangeCollectionReq) error {
	if err := s.repo.RemoveFromCart(ctx, req.UserUUID, req.DeviceUUID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetCart(ctx context.Context, userUUID string) ([]*models.Device, error) {
	col, err := s.repo.GetCart(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.CollectionToServer(col), nil
}

func (s *service) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
	if err := s.repo.RemoveDeviceFromCollections(ctx, deviceUUID); err != nil {
		return err
	}
	return nil
}
