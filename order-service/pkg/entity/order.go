package entity

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"time"
)

type OrderedDevice struct {
	Status     int32
	CreatedAt  *time.Time
	DeviceUUID string
	UserUUID   string
}

type CreateOrderReq struct {
	UserUUID    string `json:"userUUID,omitempty" validate:"required"`
	OrderUUID   string
	Status      int32
	CreatedAt   *time.Time
	DeviceUUIDs []string `json:"devices,omitempty" validate:"required"`
}

type CreateOrderReqWithDevices struct {
	UserUUID  string `json:"userUUID,omitempty" validate:"required"`
	OrderUUID string
	Status    int32
	CreatedAt time.Time
	Devices   []*pb.Device `json:"devices,omitempty" validate:"min=1"`
}

type CheckOrderRes struct {
	UserUUID  string
	OrderUUID string
	Devices   []*pb.Device
	Status    int32
	CreatedAt *time.Time
}

type UpdateOrderReq struct {
	Status    string `json:"status,omitempty" validate:"required"`
	OrderUUID string `json:"orderUUID,omitempty" validate:"required"`
}
