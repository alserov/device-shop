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

type CreateOrderReq struct {
	UserUUID  string
	OrderUUID string
	Devices   []*pb.Device
	Status    int32
	CreatedAt *time.Time
}

type CheckOrderReq struct {
}

type CheckOrderRes struct {
	UserUUID  string
	OrderUUID string
	Devices   []*pb.Device
	Status    int32
	CreatedAt *time.Time
}

type OrderedDevice struct {
	Status     int32
	CreatedAt  *time.Time
	DeviceUUID string
	UserUUID   string
}

type UpdateOrderReq struct {
}
