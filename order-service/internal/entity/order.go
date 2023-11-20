package entity

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"time"
)

type OrderAdditional struct {
	OrderUUID  string
	Status     int32
	TotalPrice float32
	CreatedAt  *time.Time
}

type OrderDevice struct {
	UUID      string
	Amount    uint32
	Status    int32
	CreatedAt *time.Time
}

type CheckOrderRes struct {
	Devices   []*pb.Device
	Status    int32
	CreatedAt *time.Time
}
