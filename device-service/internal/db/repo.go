package db

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/device-service/internal/db/models"
)

type DeviceRepo interface {
	GetActions
	AdminActions
	ChangeAmountActions
}

type GetActions interface {
	GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*models.Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]*models.Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (models.Device, error)
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]*models.Device, error)
	GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*models.Device, error)
	GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*models.Device, error)
}

type ChangeAmountActions interface {
	IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error
	DecreaseDevicesAmountTx(ctx context.Context, devices []*models.OrderDevice) (*sql.Tx, error)
}

type AdminActions interface {
	CreateDevice(context.Context, models.Device) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, models.UpdateDevice) error
}
