package db

import (
	"context"
	"database/sql"
	pb "github.com/alserov/device-shop/proto/gen"
)

type DeviceRepo interface {
	DecreaseDevicesAmount(ctx context.Context, txCh chan<- *sql.Tx, devices []*pb.OrderDevice) error
	RollbackDevices(ctx context.Context, txCh chan<- *sql.Tx, devices []*pb.Device) error
}
