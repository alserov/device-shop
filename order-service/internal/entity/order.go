package entity

import "time"

type OrderedDevice struct {
	UserUUID   string
	OrderUUID  string
	DeviceUUID string
	Amount     int32
	Status     int32
	CreatedAt  time.Time
}
