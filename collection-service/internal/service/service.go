package service

import (
	"context"

	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/db/mongo"
	"github.com/alserov/device-shop/collection-service/internal/service/models"
	"github.com/alserov/device-shop/collection-service/internal/utils/converter"

	mg "go.mongodb.org/mongo-driver/mongo"
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

func NewService(db *mg.Client, log *slog.Logger) Service {
	return &service{
		log:  log,
		db:   mongo.NewCollectionsRepo(db, log),
		conv: converter.NewServiceConverter(),
	}
}

type service struct {
	log *slog.Logger
	db  db.CollectionsRepo

	conv *converter.ServiceConverter
}

func (s *service) AddToFavourite(ctx context.Context, userUUID string, device models.Device) error {
	if err := s.db.AddToFavourite(ctx, userUUID, s.conv.Device.DeviceToRepo(device)); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveFromFavourite(ctx context.Context, req models.ChangeCollectionReq) error {
	if err := s.db.RemoveFromFavourite(ctx, req.UserUUID, req.DeviceUUID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error) {
	col, err := s.db.GetFavourite(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.CollectionToServer(col), nil
}

func (s *service) AddToCart(ctx context.Context, userUUID string, device models.Device) error {
	if err := s.db.AddToCart(ctx, userUUID, s.conv.Device.DeviceToRepo(device)); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveFromCart(ctx context.Context, req models.ChangeCollectionReq) error {
	if err := s.db.RemoveFromCart(ctx, req.UserUUID, req.DeviceUUID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetCart(ctx context.Context, userUUID string) ([]*models.Device, error) {
	col, err := s.db.GetCart(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return s.conv.Collection.CollectionToServer(col), nil
}

func (s *service) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
	if err := s.db.RemoveDeviceFromCollections(ctx, deviceUUID); err != nil {
		return err
	}
	return nil
}
