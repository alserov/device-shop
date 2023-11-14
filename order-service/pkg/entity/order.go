package entity

import "time"

type Order struct {
	UserUUID  string
	OrderUUID string
	Devices   []*Device
	Status    string
	CreatedAt time.Time
}
