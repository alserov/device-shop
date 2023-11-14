package models

type CreateOrderReq struct {
	UserUUID string    `json:"userUUID,omitempty" validate:"required"`
	Devices  []*Device `json:"devices,omitempty" validate:"required"`
}

type UpdateOrderReq struct {
}

type CheckOrderReq struct {
}
