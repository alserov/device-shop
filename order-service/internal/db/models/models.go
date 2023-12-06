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

type UpdateOrderReq struct {
	Status    string
	OrderUUID string
}

type OrderDevice struct {
	DeviceUUID string
	Amount     uint32
}

type Order struct {
	Price     float32
	CreatedAt *time.Time
	Status    int32
	UserUUID  string
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
	Status       int32
	CreatedAt    *time.Time
	OrderPrice   float32
	OrderDevices []*OrderDevice
	UserUUID     string
}
