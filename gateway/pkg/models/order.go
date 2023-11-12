package models

import "time"

type CreateOrderReq struct {
}

type OrderDevice struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       int32
	Status       string
	OrderUUID    string
	CreatedAt    *time.Time
}
