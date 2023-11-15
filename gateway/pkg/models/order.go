package models

import "github.com/alserov/device-shop/device-service/pkg/entity"

type CreateOrderReq struct {
	UserUUID string           `json:"userUUID,omitempty" validate:"required"`
	Devices  []*entity.Device `json:"devices,omitempty" validate:"required"`
}

type UpdateOrderReq struct {
}

type CheckOrderReq struct {
}
