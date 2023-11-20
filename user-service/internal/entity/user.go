package entity

import (
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	"time"
)

type User struct {
	UUID         string
	Username     string
	Password     string
	Role         string
	Email        string
	Cash         float32
	RefreshToken string
	Token        string
	CreatedAt    time.Time
}

type RepoLoginReq struct {
	RefreshToken string
	Username     string
}

type RepoLoginRes struct {
	UUID string
}

type RepoGetInfoRes struct {
	Username string
	Email    string
	UUID     string
	Cash     float32
}

type AddReq struct {
	Device   *device.Device
	UserUUID string
}

type RemoveReq struct {
	DeviceUUID string
	UserUUID   string
}

type TopUpBalanceReq struct {
	Cash     float32
	UserUUID string
}

type DebitBalanceReq struct {
	Cash     float32
	UserUUID string
}
