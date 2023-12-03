package db

import (
	"context"
)

type CollectionsRepo interface {
	AddToFavourite(ctx context.Context, userUUID string, device Device) error
	RemoveFromFavourite(ctx context.Context, userUUID string, deviceUUID string) error
	GetFavourite(ctx context.Context, userUUID string) ([]*Device, error)

	AddToCart(ctx context.Context, userUUID string, device Device) error
	RemoveFromCart(ctx context.Context, userUUID string, deviceUUID string) error
	GetCart(ctx context.Context, userUUID string) ([]*Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}
