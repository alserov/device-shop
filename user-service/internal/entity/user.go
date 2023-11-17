package entity

import (
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UUID         string
	Username     string
	Password     string
	Role         string
	Email        string
	Cash         int
	RefreshToken string
	Token        string
	CreatedAt    time.Time
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u User) CheckPassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}

type RepoLoginReq struct {
	RefreshToken string
	Username     string
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
