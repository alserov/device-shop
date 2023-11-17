package entity

import (
	"github.com/alserov/device-shop/device-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"time"
)

type Order struct {
	UserUUID  string
	OrderUUID string
	Devices   []*entity.Device
	Status    string
	CreatedAt time.Time
}

type OrderedDevice struct {
	Status     int32
	CreatedAt  *time.Time
	DeviceUUID string
	UserUUID   string
}

type CreateOrderReq struct {
	UserUUID  string `json:"userUUID,omitempty" validate:"required"`
	OrderUUID string `json:"orderUUID,omitempty" validate:"required"`
	Status    int32  `json:"status,omitempty" validate:"required"`
	CreatedAt *time.Time
	Devices   []string `json:"devices,omitempty" validate:"min=1"`
}

type CreateOrderReqWithDevices struct {
	UserUUID  string `json:"userUUID,omitempty" validate:"required"`
	OrderUUID string `json:"orderUUID,omitempty" validate:"required"`
	Status    int32  `json:"status,omitempty" validate:"required"`
	CreatedAt *time.Time
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
