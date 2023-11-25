package db

import (
	"context"
)

type AdminRepo interface {
	CreateDevice(context.Context, Device) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, UpdateDevice) error
	IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error
}

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type UpdateDevice struct {
	Title       string
	Description string
	Price       float32
	UUID        string
}
