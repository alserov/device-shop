package models

import (
	"time"
)

type CreateOrderReq struct {
	UserUUID     string
	OrderDevices []*OrderDevice
	OrderPrice   float32
}

type OrderDevice struct {
	DeviceUUID string
	Amount     uint32
}

type CreateOrderRes struct {
	OrderUUID string
}

type CheckOrderReq struct {
	OrderUUID string
}

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type CheckOrderRes struct {
	Status      string
	OrderPrice  float32
	CreatedAt   *time.Time
	DeviceUUIDs []string
}
