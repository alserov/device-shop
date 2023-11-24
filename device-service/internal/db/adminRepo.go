package db

import (
	"context"
	pb "github.com/alserov/device-shop/proto/gen"
)

type AdminRepo interface {
	CreateDevice(context.Context, *pb.Device) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, *pb.UpdateDeviceReq) error
	IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error
}
