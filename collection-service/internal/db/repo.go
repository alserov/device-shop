package db

import (
	"context"
	"github.com/alserov/device-shop/collection-service/internal/db/models"
)

type CollectionsRepo interface {
	AddToFavourite(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromFavourite(ctx context.Context, userUUID string, deviceUUID string) error
	GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error)

	AddToCart(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromCart(ctx context.Context, userUUID string, deviceUUID string) error
	GetCart(ctx context.Context, userUUID string) ([]*models.Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}
