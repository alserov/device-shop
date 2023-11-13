package models

import "time"

type CreateOrderReq struct {
	UserUUID string    `json:"userUUID,omitempty" validate:"required"`
	Devices  []*Device `json:"devices,omitempty" validate:"required"`
}

type UpdateOrderReq struct {

}

type CheckOrderReq struct {

}

type Order struct {
	UserUUID  string
	OrderUUID string
	Devices   []*Device
	Status    string
	CreatedAt time.Time
}
