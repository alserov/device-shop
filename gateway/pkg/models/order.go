package models

import "time"

type CreateOrderReq struct {
}

type OrderDevice struct {
	UserUUID           string
	DeviceUUID         string
	DeviceTitle        string
	DeviceDescription  string
	DevicePrice        float32
	DeviceManufacturer string
	OrderAmount        int32
	OrderUUID          string
	OrderStatus        string
	OrderCreatedAt     *time.Time
}
