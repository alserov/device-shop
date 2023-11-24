package db

import (
	"context"
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
)

type CollectionsRepo interface {
	AddToFavourite(ctx context.Context, userUUID string, device *pb.Device) error
	RemoveFromFavourite(ctx context.Context, req *pb.ChangeCollectionReq) error
	GetFavourite(ctx context.Context, userUUID string) ([]*device.Device, error)

	AddToCart(ctx context.Context, userUUID string, device *pb.Device) error
	RemoveFromCart(ctx context.Context, req *pb.ChangeCollectionReq) error
	GetCart(ctx context.Context, userUUID string) ([]*device.Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
}
