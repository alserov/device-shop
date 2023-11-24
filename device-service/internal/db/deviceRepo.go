package db

import (
	"context"
	pb "github.com/alserov/device-shop/proto/gen"
)

type DeviceRepo interface {
	GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*pb.Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]*pb.Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (*pb.Device, error)
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]*pb.Device, error)
	GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*pb.Device, error)
	GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*pb.Device, error)
}
