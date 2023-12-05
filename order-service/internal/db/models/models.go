package models

import "time"

type CreateOrderReq struct {
	UserUUID     string
	OrderUUID    string
	Status       int32
	OrderPrice   float32
	CreatedAt    *time.Time
	OrderDevices []*OrderDevice
}

type OrderDevice struct {
	DeviceUUID string
	Amount     uint32
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
	Status      int32
	CreatedAt   *time.Time
	OrderPrice  float32
	DeviceUUIDs []string
}
