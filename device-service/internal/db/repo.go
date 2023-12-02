package db

import (
	"context"
)

type DeviceRepo interface {
	GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]*Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (Device, error)
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]*Device, error)
	GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*Device, error)
	GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*Device, error)
}

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}
