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
	Status       string
	OrderPrice   float32
	CreatedAt    *time.Time
	OrderDevices []*OrderDevice
}

type UpdateOrderReq struct {
	Status    string
	OrderUUID string
}

type UpdateOrderRes struct {
	Status string
}

type TxResponse struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}
