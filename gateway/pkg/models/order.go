package models

import "time"

type CreateOrderReq struct {
}

type Order struct {
	UserUUID  string
	OrderUUID string
	Devices   []*Device
	Status    string
	CreatedAt time.Time
}
