package models

import (
	"time"
)

type CreateOrderReq struct {
	UserUUID     string
	OrderDevices []*OrderDevice
	OrderPrice   float32
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
	// -1 - failed (server error)
	// 0 - failed (user error)
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}

type Tx struct {
	Uuid string `json:"uuid"`
}

type TxBalanceReq struct {
	TxUUID     string  `json:"txUUID"`
	OrderPrice float32 `json:"orderPrice"`
	UserUUID   string  `json:"userUUID"`
	Status     uint32  `json:"status"`
}

type TxDeviceReq struct {
	OrderDevices []*OrderDevice
	TxUuid       string
	Status       uint32
}

type TxRequest struct {
	UserUUID     string         `json:"userUUID"`
	OrderPrice   float32        `json:"orderPrice"`
	OrderDevices []*OrderDevice `json:"orderDevices"`
}

type OrderDevice struct {
	DeviceUUID string `json:"deviceUUID"`
	Amount     uint32 `json:"amount"`
}

type DoTxBody struct {
	UserUUID     string
	OrderDevices []*OrderDevice
	OrderPrice   float32
}
