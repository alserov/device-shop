package db

import (
	"context"
	"github.com/alserov/device-shop/collection-service/internal/db/models"
)

type CollectionsRepo interface {
	CartActions
	FavouriteActions

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}

type CartActions interface {
	AddToCart(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromCart(ctx context.Context, userUUID string, deviceUUID string) error
	GetCart(ctx context.Context, userUUID string) ([]*models.Device, error)
}

type FavouriteActions interface {
	AddToFavourite(ctx context.Context, userUUID string, device models.Device) error
	RemoveFromFavourite(ctx context.Context, userUUID string, deviceUUID string) error
	GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error)
}
