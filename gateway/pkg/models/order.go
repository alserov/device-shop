package models

import "time"

type CreateOrderReq struct {
}

type Order struct {
	Status    string
	OrderUUID string
	Price     float32
	CreatedAt *time.Time
}
