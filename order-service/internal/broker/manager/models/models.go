package models

import (
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/service/models"
)

// Response from worker
// Status - indicates if the tx was successful, or not and why
// Message - if it was user error, returns error
// UUID - tx uuid
type Response struct {
	// -1 - failed (server error)
	// 0 - failed (user error)
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type BalanceReq struct {
	TxUUID     string  `json:"txUUID"`
	OrderPrice float32 `json:"orderPrice"`
	UserUUID   string  `json:"userUUID"`
	Status     uint32  `json:"status"`
}

type DeviceReq struct {
	OrderDevices []models.OrderDevice
	TxUUID       string
	Status       uint32
}

type CancelOrderTxBody struct {
	Repo      db.OrderRepo
	OrderUUID string

	OrderDevices []models.OrderDevice
}

type CreateOrderTxBody struct {
	Repo      db.OrderRepo
	Order     models.CreateOrderReq
	OrderUUID string

	UserUUID     string
	OrderDevices []models.OrderDevice
	OrderPrice   float32
}
