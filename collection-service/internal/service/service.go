package service

import (
	"context"
	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/db/mongo"
	"github.com/alserov/device-shop/collection-service/internal/utils/converter"
	mg "go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	AddToFavourite(ctx context.Context, userUUID string, device Device) error
	RemoveFromFavourite(ctx context.Context, req ChangeCollectionReq) error
	GetFavourite(ctx context.Context, userUUID string) ([]*Device, error)

	AddToCart(ctx context.Context, userUUID string, device Device) error
	RemoveFromCart(ctx context.Context, req ChangeCollectionReq) error
	GetCart(ctx context.Context, userUUID string) ([]*Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}

func NewService(db *mg.Client) Service {
	return &service{
		db: mongo.NewCollectionsRepo(db),
	}
}

type service struct {
	db db.CollectionsRepo
}

func (s *service) AddToFavourite(ctx context.Context, userUUID string, device Device) error {
	if err := s.db.AddToFavourite(ctx, userUUID, converter.ServiceDeviceToRepoStruct(device)); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveFromFavourite(ctx context.Context, req ChangeCollectionReq) error {
	if err := s.db.RemoveFromFavourite(ctx, req.UserUUID, req.DeviceUUID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetFavourite(ctx context.Context, userUUID string) ([]*Device, error) {
	fav, err := s.db.GetFavourite(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var devices []*Device
	for _, d := range fav {
		device := converter.RepoDeviceToServiceStruct(*d)
		devices = append(devices, &device)
	}

	return devices, nil
}

func (s *service) AddToCart(ctx context.Context, userUUID string, device Device) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) RemoveFromCart(ctx context.Context, req ChangeCollectionReq) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetCart(ctx context.Context, userUUID string) ([]*Device, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
	//TODO implement me
	panic("implement me")
}
